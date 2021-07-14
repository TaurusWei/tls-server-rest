package service

/**
 * @Author: WeiBingtao/13156050650@163.com
 * @Version: 1.0
 * @Description:
 * @Date: 2021/6/10 下午3:32
 */
import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/gojsonq"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"test/model"
	"unsafe"
)

//func (ns *NetSign) GenP10(socketFd int, certDN, keyLabel, keyType string) ([]byte, int) {
func genP10(socketFd int, certDN, keyLabel, keyType string) (string, int) {
	song := make(map[string]string)
	song["keyLabel"] = keyLabel
	song["certDn"] = "CN=CNCC"
	song["isCover"] = "true"
	bytesData, _ := json.Marshal(song)
	portStr := strconv.Itoa(socketFd)

	//res, err := http.Post("http://"+ns.Ip+":"+portStr+"/brilliance/netsign/genP10",
	res, err := http.Post("http://"+"47.95.204.66"+":"+portStr+"/brilliance/netsign/genP10",
		"application/json;charset=utf-8", bytes.NewBuffer([]byte(bytesData)))
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	str := (*string)(unsafe.Pointer(&content)) //转化为string,优化内存
	p10 := gojsonq.New().FromString(*str).Find("data.p10")
	return p10.(string), 0
}

func RandStringInt() string {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, serialNumberLimit)
	return serialNumber.String()
}

type Body struct {
	Port  int  `json:"port"`
	Cover bool `json:"cover"`
}

func GenP10(c *gin.Context) {
	keyType := c.Param("keyType")
	certDN := c.Query("certDN")
	body := Body{}
	err := model.GetBody(c.Request.Body, &body)
	p10, _ := genP10(body.Port, certDN, RandStringInt(), keyType)
	if err != nil {
		panic(err)
	}
	defer func() {
		recover()
	}()
	c.JSON(http.StatusOK, model.NewDefaultResponse(p10))
}
