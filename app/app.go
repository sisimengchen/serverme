package app

import (
	// "fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/sisimengchen/serverme/database"
	"github.com/sisimengchen/serverme/models"
	"github.com/sisimengchen/serverme/routes"
	// "github.com/sisimengchen/serverme/utils"
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

func RegisterDatabase(app *iris.Application) {
	database.DB.AutoMigrate(
		&models.User{},
	)
	// user := models.LoginByEmail("", "")
	// user := models.CreateUser(&models.User{
	// 	ID:       utils.GetUUID(),
	// 	Email:    "debugName1@gmail.com",
	// 	Name:     "debugUser1",
	// 	Password: "debugPassword1",
	// 	Phone:    "1888888888",
	// })
	// database.DB.Save(&models.User{
	// 	ID:       utils.GetUUID(),
	// 	Email:    "debugName@gmail.com",
	// 	Name:     "debugUser",
	// 	Password: "debugPassword",
	// })
	// user := models.GetUserById("e7806247-120f-4dc6-9884-272d384a71ca")
	// user := models.GetUserByEmail("debugName1@gmail.com")
	// database.DB.Where("email = ?", "debugName@gmail.com").First(&user)
	// fmt.Println(123)
	// fmt.Println(user)
	iris.RegisterOnInterrupt(func() {
		database.DB.Close()
	})
}

func AppInit() *iris.Application {
	// 创建iris应用
	app := iris.New()
	// 注册中间件
	RegisterMiddleware(app)
	// 注册错误处理
	RegisterErrors(app)
	// 注册数据库
	RegisterDatabase(app)
	// 注册模板 Reload(true)只需要在开发环境开启
	app.RegisterView(iris.HTML("./static", ".html").Reload(true))
	// 注册路由
	routes.RoutesInit(app)
	return app
}
