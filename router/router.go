package router

/**
 * @Author: WeiBingtao/13156050650@163.com
 * @Version: 1.0
 * @Description:
 * @Date: 2021/6/10 下午3:18
 */

import (
	"github.com/gin-contrib/pprof"
	"tls-server-rest/service"

	"github.com/gin-gonic/gin"
)

// CreateRouter 生成路由
func CreateRouter() *gin.Engine {
	router := gin.Default()
	pprof.Register(router)

	//router.POST("/test/genP10/:keyType", service.GenP10)
	router.POST("/invoke", service.Invoke)

	return router
}
