package config

/**
 * @Author: WeiBingtao/13156050650@163.com
 * @Version: 1.0
 * @Description:
 * @Date: 2021/7/15 下午6:47
 */

import (
	"github.com/spf13/viper"
	"math/rand"
	"time"
	//logger "github.com/sirupsen/logrus"
	logger "tls-server-rest/common/log"
)

// 变量默认配置表
const (
	TOKEN_NAME      = "X-Auth-Token" // request header中token的默认名字
	TOKEN_NOT_VALID = 0              // token无效
	TOKEN_VALID     = 1              // token有效
	TOKEN_EXPIRED   = 2              // token过期
)

// 日志配置
type LogConfig struct {
	Formatter string `yaml:"formatter"`
	Level     string `yaml:"level"`
}

// 从配置文件的路径读取配置
func InitConfig(files []string) error {
	logger.Info("加载配置文件")
	for index, file := range files {
		viper.SetConfigFile(file)
		if 0 == index {
			if err := viper.ReadInConfig(); err != nil {
				if _, ok := err.(viper.ConfigFileNotFoundError); ok {
					logger.Error("No such config file, please check the config file path!")
				} else {
					logger.Errorf("Read config file error，Err:%s", err.Error())
				}
				return err
			}
		} else {
			if err := viper.MergeInConfig(); err != nil {
				if _, ok := err.(viper.ConfigFileNotFoundError); ok {
					logger.Error("No such config file, please check the config file path!")
				} else {
					logger.Errorf("Read config file error，Err:%s", err.Error())
				}
				return err
			}
		}
	}
	return nil
}

// 获取监听地址
func GetRestfulListenAddress() string {
	return viper.GetString("server.restful.listenAddress")
}

// 获取数据库连接的解密密钥
func GetPostgresqlConnectionKey() string {
	return viper.GetString("db.postgresql.key")
}

// 获取postgresql的数据库连接
func GetPostgresqlConnection() string {
	return viper.GetString("db.postgresql.connection")
}

// 获取mysql的数据库连接
func GetMysqlConnection() string {
	return viper.GetString("db.mysql.connection")
}

// 获取报文推送锁名称
func GetMsgForwardLockName() string {
	return viper.GetString("server.lockName.MsgForwardLock")
}

// 监听报文下链入库锁
func GetListenMsgLockName() string {
	return viper.GetString("server.lockName.ListenMsgLock")
}

// 获取token过期时间
func GetTokenExpireDuration() int64 {
	return viper.GetInt64("server.token.tokenExpireDuration")
}

// 获取token名字
func GetTokenName() string {
	return viper.GetString("server.token.tokenName")
}

// 获取日志配置
func GetLogConfig() *LogConfig {
	cfg := &LogConfig{
		Formatter: viper.GetString("server.logging.formatter"),
		Level:     viper.GetString("server.logging.level"),
	}

	return cfg
}

// 获取同步更新host的间隔
func GetHostUpdateInternal() time.Duration {
	duration := viper.GetDuration("server.docker.hostUpdateInternal")
	if duration == 0 {
		duration = time.Minute * 1
	}
	return duration
}

// 是否同步开启host文件
func GetHostUpdateEnabled() bool {
	return viper.GetBool("server.docker.enabled")
}

// rabbitmq是否开启 true: 开启 false: 未开启
func GetRabbitMqEnabled() bool {
	return viper.GetBool("server.rabbitmq.enabled")
}

// rabbitmq加密私钥
func GetRabbitMqKey() string {
	return viper.GetString("server.rabbitmq.key")
}

type ConnStru struct {
	UserName     string //用户名
	NetworkName  string //网络名
	ChannelName  string //通道名
	ContractName string //合约名
	EventType    int    //事件类型
	EventFilter  string //事件名
}

//
func GetConnStru() (conn *ConnStru) {

	conn = new(ConnStru)

	conn.UserName = viper.GetString("conn.userName")
	conn.NetworkName = viper.GetString("conn.networkName")
	conn.ChannelName = viper.GetString("conn.channelName")
	conn.ContractName = viper.GetString("conn.contractName")
	conn.EventType = viper.GetInt("conn.eventType")
	conn.EventFilter = viper.GetString("conn.eventFilter")

	return
}

// 获取http client超时时间
func GetHttpTimeout() int64 {
	return viper.GetInt64("http_client.timeout")
}

// 获取cows/lcfs post地址
func GetHttpUrl(sys string) []string {
	ret := viper.GetStringSlice("http_client." + sys + ".addr")
	if 0 == len(ret) {
		return ret
	}
	//增加随机数，返回随机组合地址
	n := rand.Intn(len(ret))
	ret = append(ret[n:], ret[:n]...)
	logger.Debugf("http_client.%s.addr:随机数:[%d]地址列表:[%v]", sys, n, ret)
	return ret
}

// 获取cows/connverfication地址
func GetCOWSConnverfication() []string {
	ret := viper.GetStringSlice("http_client.cows.connverfication")
	if 0 == len(ret) {
		return ret
	}
	//增加随机数，返回随机组合地址
	n := rand.Intn(len(ret))
	ret = append(ret[n:], ret[:n]...)
	logger.Debugf("http_client.cows.connverfication:随机数:[%d]地址列表:[%v]", n, ret)
	return ret
}

//获取CBAS部署的所在集群中心地址，如北京(BJ)或上海(SH)
func GetClusterAddr() string {
	ret := viper.GetString("ClusterAddr.centers")
	return ret
}

func GetInvokeTimeoutWait() time.Duration {
	duration := viper.GetDuration("server.wait.invokeTimeout")
	if duration == 0 {
		duration = time.Second * 30
	}
	return duration
}

func GetPoolGoroutines() uint {
	ret := viper.GetInt("server.pool.GoroutinesNum")
	//默认为50
	if ret <= 0 {
		ret = 50
	}
	result := uint(ret)
	logger.Debugf("PoolGoroutines大小为：%d", result)
	return result
}

/**获取当前中心的运行状态
ocinfo.yaml:
STATUS:
  BJ: ACTV
  SH: DOWN
**/
func GetClusterSts(cluster string) string {
	ret := viper.GetString("STATUS." + cluster)
	return ret
}

//获取CBAS实例名
func GetHostName() string {
	ret := viper.GetString("base.hostname")
	return ret
}

/**
  获取 tls server cert
*/
func GetTlsServerCert() string {
	ret := viper.GetString("tls.server.cert")
	return ret
}

/**
  获取 tls server key
*/
func GetTlsServerKey() string {
	ret := viper.GetString("tls.server.key")
	return ret
}

/**
  获取 server ecert  : 签名证书
*/
func GetServerCert() string {
	ret := viper.GetString("server.cert")
	return ret
}
