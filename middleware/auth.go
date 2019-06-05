package middleware

import (
	"github.com/kataras/iris"
	"github.com/sisimengchen/serverme/models"
	"github.com/sisimengchen/serverme/utils"
	"github.com/sisimengchen/serverme/controllers"
)

func Auth(ctx iris.Context) {
	// 判断是否是需要做权限验证的path
	if utils.IsAuthRequest(ctx) {
		id := utils.GetSecureCookie(ctx, "fess")
		if len(id) > 0 {
			user, err := models.GetUser(&models.Users{ ID: id })
			if err == nil {
				ctx.Values().Set("contextUser", *user)
			}
		} else {
			if !utils.IsApiRequest(ctx) {
				ctx.Redirect("/page/login")
			} else {
				ctx.JSON(controllers.ResponseResource(401, "unlogin", nil))
			}
			return
		}
	}
	ctx.Next()
}
