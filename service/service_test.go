package service

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"testing"
)

/**
 * @Author: WeiBingtao/13156050650@163.com
 * @Version: 1.0
 * @Description:
 * @Date: 2021/8/6 下午3:17
 */
func Test_https_get(t *testing.T) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Get("https://47.95.204.66:34997/brilliance//netsign/reset")
	fmt.Println(res)
	fmt.Println(err)
}
