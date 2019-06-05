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
	// 静态资源
	app.StaticWeb("/public", "./public")
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
	api := app.Party("/api")
	{
		api.Post("/login", controllers.Login)
		api.Get("/logout", controllers.Logout)
		api.Post("/reg", controllers.CreateUser)

		auth := api.Party("/auth", middleware.Auth)
		{
			auth.Get("/user", controllers.GetUser)
			auth.Post("/passreset", controllers.UpdateUserPassword)
			auth.Post("/updateuser", controllers.UpdateUser)
			auth.Post("/updateavatar", controllers.UpdateAvatar)
			
			common := api.Party("/common")
			{
				common.Post("/upload", controllers.Upload)
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
		}		
	}
}
