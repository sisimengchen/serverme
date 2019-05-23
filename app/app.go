package app

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/sisimengchen/serverme/routes"
	// "github.com/sisimengchen/serverme/middleware"
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

func RegisterMiddleware(app *iris.Application) {
	// 恢复中间件
	app.Use(recover.New())
	// 日志中间件
	app.Use(logger.New())
	// // 鉴权中间件
	// app.UseGlobal(middleware.Auth)
}

func AppInit() *iris.Application {
	// 创建iris应用
	app := iris.New()
	// 注册中间件
	RegisterMiddleware(app)
	// 注册错误处理
	RegisterErrors(app)
	// 注册模板 Reload(true)只需要在开发环境开启
	app.RegisterView(iris.HTML("./static", ".html").Reload(true))
	// 注册路由
	routes.RoutesInit(app)
	return app
}
