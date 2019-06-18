package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sisimengchen/serverme/controllers"
	"github.com/sisimengchen/serverme/middleware"
	"net/http"
)

func Init(router *gin.Engine) {

	router.StaticFile("/favicon.ico", "static/favicon.ico")

	router.Static("/static", "static")

	router.Static("/public", "public")

	router.Static("/uploads", "uploads")

	router.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404.html", nil)
	})

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
	api := router.Group("/api", middleware.Cors())
	{
		api.POST("/login", controllers.Login)
		api.GET("/logout", controllers.Logout)
		api.POST("/reg", controllers.CreateUser)
		api.GET("/hot", controllers.GetHotBooks)

		auth := api.Group("/auth", middleware.Auth() /*, middleware.Casbin()*/)
		{
			auth.GET("/user", controllers.GetUser)
			auth.GET("/userroles", controllers.GetUserWithRoles)
			auth.GET("/user/roles", controllers.GetUserRoles)
			auth.POST("/passreset", controllers.UpdateUserPassword)
			auth.POST("/updateuser", controllers.UpdateUser)
			auth.POST("/updateavatar", controllers.UpdateAvatar)

			common := auth.Group("/common")
			{
				common.POST("/upload", controllers.Upload)
			}

			role := auth.Group("/role")
			{
				role.POST("/add", controllers.CreateRole)
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
	var secrets = gin.H{
		"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
		"austin": gin.H{"email": "austin@example.com", "phone": "666"},
		"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
	}
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))
	{
		authorized.GET("/secrets", func(c *gin.Context) {
			// get user, it was set by the BasicAuth middleware
			user := c.MustGet(gin.AuthUserKey).(string)
			if secret, ok := secrets[user]; ok {
				c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
			} else {
				c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
			}
		})
	}
}
