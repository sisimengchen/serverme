package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sisimengchen/serverme/controllers"
	"github.com/sisimengchen/serverme/middleware"
	"net/http"
)

func RoutesInit(router *gin.Engine) {

	router.StaticFile("/favicon.ico", "./static/favicon.ico")

	router.Static("/static", "./static")

	router.Static("/public", "./public")

	router.Static("/uploads", "./uploads")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	page := router.Group("/page")
	{
		page.GET("/", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index.html", nil)
		})
		page.GET("/:any", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index.html", nil)
		})
	}

	api := router.Group("/api")
	{
		api.POST("/login", controllers.Login)
		api.GET("/logout", controllers.Logout)
		api.POST("/reg", controllers.CreateUser)

		auth := api.Group("/auth", middleware.Auth())
		{
			auth.GET("/user", controllers.GetUser)
			auth.POST("/passreset", controllers.UpdateUserPassword)
			auth.POST("/updateuser", controllers.UpdateUser)
			auth.POST("/updateavatar", controllers.UpdateAvatar)

			common := auth.Group("/common")
			{
				common.POST("/upload", controllers.Upload)
			}

			bookcatalog := auth.Group("/bookcatalog")
			{
				bookcatalog.POST("/add", controllers.CreateBookCatalog)
				bookcatalog.GET("/get", controllers.GetBookCatalogByID)
				bookcatalog.GET("/getall", controllers.GetBookCatalogs)
			}

			book := auth.Group("/book")
			{
				book.POST("/add", controllers.CreateBook)
				book.GET("/get", controllers.GetBookByID)
				book.GET("/getbyname", controllers.GetBooksByName)
				book.GET("/getbycatalog", controllers.GetBooksByCatalogId)
				book.GET("/getall", controllers.GetBooks)
			}

			chapter := auth.Group("/chapter")
			{
				chapter.POST("/add", controllers.CreateChapter)
				chapter.POST("/setpath", controllers.SetChapterPath)
				chapter.GET("/get", controllers.GetChapterByID)
				chapter.GET("/gets", controllers.GetChaptersByBookId)
				chapter.GET("/getcontent", controllers.GetChapterContent)
			}
		}
	}
}
