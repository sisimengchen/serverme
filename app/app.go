package app

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sisimengchen/serverme/routes"
)

func Init() *gin.Engine {
	// gin.DebugMode gin.ReleaseMode gin.TestMode
	gin.SetMode(gin.DebugMode)
	// 创建iris应用
	router := gin.New()
	// 恢复中间件
	router.Use(gin.Recovery())
	// 日志中间件
	router.Use(gin.Logger())
	// gzip中间件，也可以统一在ng层做设置
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	// 注册模板
	router.LoadHTMLGlob("views/*")
	// 注册路由
	routes.Init(router)

	return router
}
