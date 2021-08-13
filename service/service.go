package service

/**
 * @Author: WeiBingtao/13156050650@163.com
 * @Version: 1.0
 * @Description:
 * @Date: 2021/6/10 下午3:32
 */
import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric/bccsp"
	"github.com/pkg/errors"

	"github.com/hyperledger/fabric/bccsp/cncc"
	factory "github.com/hyperledger/fabric/bccsp/factory"
	"github.com/hyperledger/fabric/bccsp/verifier"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"tls-server-rest/backend/dao"
	config1 "tls-server-rest/common/config"
	"tls-server-rest/common/global"
	logger "tls-server-rest/common/log"
	"tls-server-rest/common/util"
	"tls-server-rest/model"
)

var (
	poolMutex    sync.RWMutex // 读写锁
	configPath   string       // 配置路径 从配置路径初始化fabricsdk
	cryptoConfig = &CryptoConfig{}
	Client       *http.Client // https请求客户端
)

type CryptoConfig struct {
	Provider  string
	netsigns  cncc.NetSignConfig
	networkId string
}

func NewService(config string) (bccsp.BCCSP, error) {
	configPath = config
	//获取网络信息
	// todo networkName
	networks, err := dao.GetNetworkByName(config1.GetConnStru().NetworkName)
	if err != nil || len(networks) == 0 {
		logger.Errorf("查询网络信息失败: %s", err.Error())
		return nil, fmt.Errorf("查询网络信息失败: %s", err.Error())
	}
	network := networks[0]
	//cryptoConfig只初始化一次，每次复用
	// todo networkName
	signServers, err := dao.GetAllSignServiceInfo(config1.GetConnStru().NetworkName)
	if err != nil {
		logger.Errorf("查询签名服务器信息失败: %s", err.Error())
		return nil, fmt.Errorf("查询签名服务器信息失败: %s", err.Error())
	}
	// todo ClusterAddr.center
	ip, port, password := combinNetsignsUsedForBaas(signServers, config1.GetClusterAddr())
	var signServerInfo []cncc.NetSignConfig
	signServerInfo0 := cncc.NetSignConfig{
		Ip:     ip,
		Port:   port,
		Passwd: password,
	}
	signServerInfo = append(signServerInfo, signServerInfo0)
	cryptoConfig.Provider = network.CryptoProvider
	cryptoConfig.netsigns = signServerInfo[0]
	cryptoConfig.networkId = strconv.Itoa(network.Id)

	opts := factory.GetDefaultOpts()
	opts.ProviderName = network.CryptoProvider
	opts.CNCC_GMOpts.Ip = cryptoConfig.netsigns.Ip
	opts.CNCC_GMOpts.Port = cryptoConfig.netsigns.Port
	opts.CNCC_GMOpts.Password = cryptoConfig.netsigns.Passwd
	opts.CNCC_GMOpts.NetWorkId = cryptoConfig.networkId

	csp, err := (&factory.CNCC_GMFactory{}).Get(opts)
	if err != nil {
		logger.Errorf("获取 Bccsp 实例失败：%s", err.Error())
		return nil, fmt.Errorf("获取Bccsp实例失败：%s", err.Error())
	}
	factory.SetBCCSP(network.CryptoProvider, csp)

	verifier, err := verifier.New(csp, nil)
	if err != nil {
		logger.Errorf("初始化 verifier ：%s", err.Error())
		return nil, fmt.Errorf("初始化 verifier ：%s", err.Error())
	}
	global.Verifier = verifier

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	Client = &http.Client{Transport: tr}
	return csp, nil
}
func combinNetsignsUsedForBaas(netsigns []dao.SignService, dc string) (string, string, string) {
	var ip, port, ps []string
	var dcip, dcport, dcps []string
	for _, net := range netsigns {
		if net.DataCenter == dc {
			dcip = append(dcip, net.IP)
			dcport = append(dcport, strconv.Itoa(net.Port))
			dcps = append(dcps, net.Password)
		} else {
			ip = append(ip, net.IP)
			port = append(port, strconv.Itoa(net.Port))
			ps = append(ps, net.Password)
		}
	}
	ips, ports, pss := strings.Join(ip, ","), strings.Join(port, ","), strings.Join(ps, ",")
	dcips, dcports, dcpss := strings.Join(dcip, ","), strings.Join(dcport, ","), strings.Join(dcps, ",")
	cnccip, cnccport, cnccps := []string{dcips, ips}, []string{dcports, ports}, []string{dcpss, pss}

	return combinStr(cnccip, ";"), combinStr(cnccport, ";"), combinStr(cnccps, ";")
}

func combinStr(strs []string, sep string) string {
	return strings.Trim(strings.Join(strs, sep), sep)
}

func Invoke(c *gin.Context) {

	request := model.Envelope{}
	err := model.GetBody(c.Request.Body, &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	result, err := invoke(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}
func invoke(request model.Envelope) (*model.Envelope, error) {
	verify, err := util.Verify(request.Sig, request.Data, request.Certificate)
	if err != nil || !verify {
		logger.Errorf("The request data cannot be verified: %s", err.Error())
		return nil, errors.WithMessagef(err, "The request data cannot be verified")
	}
	requestData := request.Data
	queryBaseInfo := model.QueryBaseInfo{}
	err = json.Unmarshal(requestData, &queryBaseInfo)
	if err != nil {
		logger.Errorf("Unmarshal QueryBaseInfo error: %s", err.Error())
		return nil, errors.WithMessagef(err, "Unmarshal QueryBaseInfo error")
	}
	logger.Infof("request info: %v", queryBaseInfo)
	bytesData, _ := json.Marshal(queryBaseInfo.Params)
	var res *http.Response
	// todo  add get method
	if strings.HasPrefix(queryBaseInfo.Url, "https://") {
		if strings.EqualFold(queryBaseInfo.Method, "post") {
			res, err = Client.Post(queryBaseInfo.Url, "application/json;charset=utf-8", bytes.NewBuffer(bytesData))
		} else {
			res, err = Client.Get(queryBaseInfo.Url)
		}
	} else {
		if strings.EqualFold(queryBaseInfo.Method, "post") {
			res, err = http.Post(queryBaseInfo.Url, "application/json;charset=utf-8", bytes.NewBuffer(bytesData))
		} else {
			res, err = http.Get(queryBaseInfo.Url)
		}
	}
	if err != nil {
		logger.Errorf("query data error: %s", err.Error())
		return nil, errors.WithMessagef(err, "query data error")
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)

	if err != nil {
		logger.Errorf("read response body error: %s", err.Error())
		return nil, errors.WithMessagef(err, "read response body error")
	}
	sig, err := util.Sign(content)
	if err != nil {
		logger.Errorf("sign error: %s", err.Error())
		return nil, err
	}
	// todo sign cert
	certBytes, err := ioutil.ReadFile(config1.GetServerCert())
	if err != nil {
		logger.Errorf("读取签名证书失败, path = %s , error: ", config1.GetServerCert(), err.Error())
		return nil, errors.WithMessagef(err, "读取签名证书失败, path = %s ", config1.GetServerCert())
	}
	envelope := &model.Envelope{
		Data:        content,
		Sig:         sig,
		Certificate: certBytes,
	}
	return envelope, nil
}
