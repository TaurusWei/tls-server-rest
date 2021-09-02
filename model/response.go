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
	SUCCESS                = 200
	ERROR                  = 500
)

var codeMsg = map[int]string{
	SUCCESS:            "ok",
	ERROR:              "fail",
	PARAM_VALUE_WRONG:  "参数取值不对",
	PARAM_FORMAT_WRONG: "参数格式不对",
	PARAM_REQUIRED:     "必输参数不能为空",
	PARAM_TOO_LONG:     "参数超长",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}

type Response struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data"`
}

func NewResponse(code int, data interface{}) Response {
	switch data.(type) {
	case string:
		fmt.Printf("Response: %s", data.(string))
		return Response{Code: code, Msg: data.(string)}
	default:
		jData, err := json.Marshal(data)
		if err != nil {
			return Response{Code: http.StatusInternalServerError}
		}
		fmt.Printf("Response: %s", string(jData))
		return Response{Code: code, Data: jData}
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
