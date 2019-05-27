package controllers

import (
	// "fmt"
	"github.com/kataras/iris"
	"github.com/sisimengchen/serverme/models"
	"github.com/sisimengchen/serverme/utils"
)

func UserLogin(ctx iris.Context) {
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")
	if email == "" || password == "" {
		ctx.JSON(iris.Map{
			"message": "fail",
		})
	} else {
		user := models.LoginByEmail(email, password)
		if user != nil {
			utils.SetSecureCookie(ctx, "fess", user.ID)
			ctx.JSON(iris.Map{
				"message": "ok",
			})
		} else {
			ctx.JSON(iris.Map{
				"message": "fail",
			})
		}
	}
}

func UserLogout(ctx iris.Context) {
	utils.RemoveCookie(ctx, "fess")
	utils.RemoveCookie(ctx, "fess.sig")
	ctx.JSON(iris.Map{
		"message": "ok",
	})
}

func UserInfo(ctx iris.Context) {
	user := ctx.Values().Get("user").(models.User)
	ctx.JSON(iris.Map{
		"message": "ok",
		"email":   user.Email,
		"phone":   user.Phone,
	})
}

func UserCreate(ctx iris.Context) {
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")
	if email == "" || password == "" {
		ctx.JSON(iris.Map{
			"message": "fail",
		})
		return
	}
	user := models.GetUserByEmail(email)
	if user != nil {
		ctx.JSON(iris.Map{
			"message": "another",
		})
		return
	}
	user = &models.User{ID: utils.GetUUID(), Email: email, Password: password}
	user = models.CreateUser(user)
	if user == nil {
		ctx.JSON(iris.Map{
			"message": "fail",
		})
		return
	}
	UserLogin(ctx)
}
