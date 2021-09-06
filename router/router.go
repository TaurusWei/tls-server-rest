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
	//swaggerFiles "github.com/swaggo/files"
	"github.com/gin-gonic/gin"
)

// CreateRouter 生成路由
func CreateRouter() *gin.Engine {
	router := gin.Default()
	pprof.Register(router)
	testGroup := router.Group("/test/")
	testGroup.POST("/invokeTest", service.InvokeTest)

	oracleGroup := router.Group("/oracle/")
	// swagger
	//oracleGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//router.POST("/test/genP10/:keyType", service.GenP10)
	oracleGroup.POST("/invoke", service.Invoke)

	return router
}
