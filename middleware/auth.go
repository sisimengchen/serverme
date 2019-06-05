package middleware

import (
	"github.com/kataras/iris"
	"github.com/sisimengchen/serverme/models"
	"github.com/sisimengchen/serverme/utils"
)

// var (
// 	debugUser = models.Users{
// 		ID:       utils.GetUUID(),
// 		Email:    "debuguser@gmail.com",
// 		Name:     "debugusername",
// 	}
// )

func Auth(ctx iris.Context) {
	// 判断是否是需要做权限验证的path
	if utils.IsAuthRequest(ctx) {
		fess := utils.GetSecureCookie(ctx, "fess")
		if len(fess) > 0 {
			user, err := models.GetUserByID(fess)
			if err == nil {
				ctx.Values().Set("contextUser", *user)
			}
		} else {
			if !utils.IsApiRequest(ctx) {
				ctx.Redirect("/login")
			}
		}
	}
	ctx.Next()
}
