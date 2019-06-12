package app

import (
	// "fmt"
	"github.com/kataras/iris"
	// "github.com/kataras/iris/middleware/logger"
	// "github.com/kataras/iris/middleware/recover"
	"github.com/gin-gonic/gin"
	"github.com/sisimengchen/serverme/database"
	"github.com/sisimengchen/serverme/models"
	"github.com/sisimengchen/serverme/routes"
)

const notFoundMessage = "404 Not Found"
const internalServerErrorMessage = "Oups something went wrong, try again"
const notFoundHTML = "<h1>" + notFoundMessage + "</h1>"
const internalServerErrorHTML = "<h1>" + internalServerErrorMessage + "</h1>"

func RegisterErrors(app *iris.Application) {
	// 404处理
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.HTML(notFoundHTML)
	})
	// 500处理
	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.HTML(internalServerErrorHTML)
	})
}

func RegisterMiddleware(engine *gin.Engine) {
	// 恢复中间件
	engine.Use(gin.Logger())
	// 日志中间件
	engine.Use(gin.Recovery())
	// // 鉴权中间件
	// app.UseGlobal(middleware.Auth)
}

func RegisterDatabase(app *iris.Application) {
	database.DB.AutoMigrate(
		&models.Users{},
		&models.BookCatalogs{},
		&models.Books{},
		&models.Chapters{},
	)
	iris.RegisterOnInterrupt(func() {
		database.DB.Close()
	})
}

func EngineRouter() *gin.Engine {
	// 创建iris应用
	router := gin.New()
	// 恢复中间件
	router.Use(gin.Logger())
	// 日志中间件
	router.Use(gin.Recovery())
	// 注册中间件
	// RegisterMiddleware(engine)
	// 注册错误处理
	// RegisterErrors(app)
	// 注册数据库
	// RegisterDatabase(app)
	// 注册模板 Reload(true)只需要在开发环境开启
	// app.RegisterView(iris.HTML("./static", ".html").Reload(true))
	router.LoadHTMLGlob("./static/*")
	// 注册路由
	routes.RoutesInit(router)

	return router
}
