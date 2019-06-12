package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sisimengchen/serverme/controllers"
	"github.com/sisimengchen/serverme/models"
	"github.com/sisimengchen/serverme/utils"
	"net/http"
)

// 重定向的url之后要做成可配置
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, _ := ctx.Cookie("fess")
		if len(id) > 0 {
			user, err := models.GetUser(&models.Users{ID: id})
			if err == nil {
				ctx.Set("contextUser", *user)
				ctx.Next()
				return
			}
		} else {
			if !utils.IsApiRequest(ctx) {
				ctx.Redirect(http.StatusFound, "/page/login")
			} else {
				ctx.JSON(controllers.ResponseResource(http.StatusUnauthorized, "unlogin", nil))
			}
			ctx.Abort()
			return
		}
	}
}
