package routes

import (
	"github.com/kataras/iris"
	"github.com/sisimengchen/serverme/controllers"
	"github.com/sisimengchen/serverme/middleware"
	// "github.com/dgrijalva/jwt-go"
    // jwtmiddleware "github.com/iris-contrib/middleware/jwt"
)

func RoutesInit(app *iris.Application) {
	// favicon
	app.Favicon("./static/favicon.ico")
	// 静态资源
	app.StaticWeb("/static", "./static")
	// 上传资源
	// app.StaticWeb("/uploads", "./uploads")
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
		api.Post("/login", controllers.UserLogin)
		api.Get("/logout", controllers.UserLogout)
		api.Post("/reg", controllers.UserCreate)
		api.Post("/auth/passreset", controllers.UserPasswordUpdate)
		api.Get("/auth/user", controllers.UserInfo)

		common := api.Party("/auth/common")
		{
			common.Post("/upload", controllers.Upload)
		}

		configure := api.Party("/auth/configure")
		{
			configure.Post("/add", controllers.ConfigureCreate)
			configure.Get("/get", controllers.ConfigureGet)
			configure.Get("/getbytopic", controllers.ConfiguresGetByTopic)
		}

		book := api.Party("/auth/book")
		{
			book.Post("/add", controllers.CreateBook)
			book.Get("/get", controllers.GetBookByID)
			book.Get("/getbyname", controllers.GetBooksByName)
			book.Get("/getbycatalog", controllers.GetBooksByCatalogId)
			book.Get("/getall", controllers.GetBooks)
		}
		
		chapter := api.Party("/auth/chapter")
		{
			chapter.Post("/add", controllers.CreateChapter)
			chapter.Post("/setpath", controllers.SetChapterPath)
			chapter.Get("/get", controllers.GetChapterByID)
			chapter.Get("/gets", controllers.GetChaptersByBookId)
			chapter.Get("/getcontent", controllers.GetChapterContent)
		}
		
		bookcatalog := api.Party("/auth/bookcatalog")
		{
			bookcatalog.Post("/add", controllers.CreateBookCatalog)
			bookcatalog.Get("/get", controllers.GetBookCatalogByID)
			bookcatalog.Get("/getall", controllers.GetBookCatalogs)
		}
		// api.Get("/jwt", func(ctx iris.Context) {
		// 	token := ctx.Values().Get("jwt").(*jwt.Token)
		// 	ctx.Writef("This is an authenticated request\n")
        //     ctx.Writef("Claim content:\n")
        //     //可以了解一下token的数据结构
        //     ctx.Writef("%s", token)
		// })
	}
}
