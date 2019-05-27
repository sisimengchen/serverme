package middleware

import (
	"github.com/kataras/iris"
	"github.com/sisimengchen/serverme/models"
	"github.com/sisimengchen/serverme/utils"
)

var (
	debugUser = models.User{
		ID:       utils.GetUUID(),
		Email:    "debuguser@gmail.com",
		Name:     "debugusername",
		Password: "debuguserpassword",
	}
)

func Auth(ctx iris.Context) {
	// 判断是否是需要做权限验证的path
	if utils.IsAuthRequest(ctx) {
		// 如果开启了debugmode模式
		debugmode := utils.GetCookie(ctx, "debugmode")
		// 如果开启了debug模式
		if debugmode == "debugmode" {
			ctx.Values().Set("user", debugUser)
			ctx.Next()
		} else {
			fess := utils.GetSecureCookie(ctx, "fess")
			if fess != "" {
				user := models.GetUserById(fess)
				if user != nil {
					ctx.Values().Set("user", *user)
					ctx.Next()
				} else {
					if utils.IsApiRequest(ctx) {
						ctx.JSON(iris.Map{
							"message": "nologin",
						})
					} else {
						ctx.Redirect("/login")
					}
				}
			} else {
				if utils.IsApiRequest(ctx) {
					ctx.JSON(iris.Map{
						"message": "nologin",
					})
				} else {
					ctx.Redirect("/login")
				}
			}
		}
	} else {
		ctx.Next()
	}

}
