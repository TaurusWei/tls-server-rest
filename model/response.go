package model

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/**
 * @Author: WeiBingtao/13156050650@163.com
 * @Version: 1.0
 * @Description:
 * @Date: 2021/6/10 下午3:52
 */
const (
	PARAM_VALUE_WRONG  = 10003 // 参数取值不对
	PARAM_FORMAT_WRONG = 10004 // 参数格式不对
	PARAM_REQUIRED     = 10005 // 必输参数
	PARAM_TOO_LONG     = 10006 // 参数超长

	PARAM_VALUE_WRONG_MSG  = "参数取值不对"
	PARAM_FORMAT_WRONG_MSG = "参数格式不对"
	PARAM_REQUIRED_MSG     = "必输参数不能为空"
	PARAM_TOO_LONG_MSG     = "参数超长"
)

type Response interface {
}

type DefaultResponse struct {
	Code int             `json:"code"`
	Data json.RawMessage `json:"data"`
}

type StringResponse struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

func NewResponse(code int, data interface{}) Response {
	switch data.(type) {
	case string:
		fmt.Printf("Response: %s", data.(string))
		return StringResponse{Code: code, Data: data.(string)}
	default:
		jData, err := json.Marshal(data)
		if err != nil {
			return DefaultResponse{Code: http.StatusInternalServerError}
		}
		fmt.Printf("Response: %s", string(jData))
		return DefaultResponse{Code: code, Data: jData}
	}
}

func NewDefaultResponse(data interface{}) Response {
	return NewResponse(http.StatusOK, data)
}

func NewErrorResponse(err error) Response {
	return NewResponse(http.StatusInternalServerError, err.Error())
}

func NewUnauthorizedResponse() Response {
	return NewResponse(http.StatusUnauthorized, "未经授权的访问1")
}
