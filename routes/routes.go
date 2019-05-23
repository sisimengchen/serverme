package routes

import (
	"github.com/kataras/iris"
	"github.com/sisimengchen/serverme/middleware"
	"github.com/sisimengchen/serverme/models"
	// "github.com/dgrijalva/jwt-go"
)

func RoutesInit(app *iris.Application) {
	// favicon
	app.Favicon("./static/favicon.ico")
	// 静态资源
	app.StaticWeb("/static", "./static")
	// 单页路由
	app.Get("/", middleware.Auth, func(ctx iris.Context) {
		ctx.View("index.html")
	})
	// page路由 (视图层)
	page := app.Party("/page", middleware.Auth)
	{
		page.Get("/*", func(ctx iris.Context) {
			ctx.View("index.html")
		})
	}
	// api路由（接口层）
	api := app.Party("/api", middleware.Auth, middleware.Cros())
	{
		api.Get("/auth/ping", func(ctx iris.Context) {
			// user := ctx.Values().Get("jwt").(*jwt.Token)
			user := ctx.Values().Get("user").(models.User)
			ctx.JSON(iris.Map{
				"message":  "apipong",
				"username": user.Username,
				// "user": user,
			})
		})
	}
}
