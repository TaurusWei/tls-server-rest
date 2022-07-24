package router

/**
 * @Author: WeiBingtao/13156050650@163.com
 * @Version: 1.0
 * @Description:
 * @Date: 2021/6/10 下午3:18
 */

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "tls-server-rest/docs"
	"tls-server-rest/service"
)

// CreateRouter 生成路由
func CreateRouter() *gin.Engine {
	router := gin.Default()
	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//pprof.Register(router)
	testGroup := router.Group("/test/")
	testGroup.POST("/invokeTest", service.InvokeTest)

	oracleGroup := router.Group("/oracle/")
	//oracleGroup.POST("/genP10/:keyType", service.GenP10)
	oracleGroup.POST("/invoke", service.Invoke)

	return router
}
