package model

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

/**
 * @Author: WeiBingtao/13156050650@163.com
 * @Version: 1.0
 * @Description:
 * @Date: 2021/6/10 下午3:46
 */
/*
	业务请求结构体
*/
type Envelope struct {
	Data        []byte `json:"data"`        // 业务合约请求数据, 对应 QueryBaseInfo
	Sig         []byte `json:"sig"`         // 签名值
	Certificate []byte `json:"certificate"` // 证书
}

/*
  业务合约请求数据  结构体
*/
type QueryBaseInfo struct {
	ContractName string                 `json:contractName,omitempty"` // 应用合约名字
	Method       string                 `json:"method,omitempty"`      // 请求的方法
	Params       map[string]interface{} `json:"params,omitempty"`      // 请求的数据
	SourceUrl    string                 `json:sourceUrl,omitempty"`    // 数据源地址
}

func GetBody(body io.ReadCloser, v interface{}) error {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}
