package controllers

import (
	"github.com/kataras/iris"
	"github.com/sisimengchen/serverme/models"
	"github.com/sisimengchen/serverme/utils"
)

func UserLogin(ctx iris.Context) {
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")
	user, err := models.LoginByEmail(email, password)
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		utils.SetSecureCookie(ctx, "fess", user.ID)
		ctx.JSON(ResponseResource(200, "ok", nil))
	}
}

func UserLogout(ctx iris.Context) {
	utils.RemoveCookie(ctx, "fess")
	utils.RemoveCookie(ctx, "fess.sig")
	ctx.JSON(ResponseResource(200, "ok", nil))
}

func UserInfo(ctx iris.Context) {
	contextUser, err := GetContextUser(ctx)
	if err != nil {
		ctx.JSON(ResponseResource(401, err.Error(), nil))
		return
	}
	ctx.JSON(ResponseResource(200, "ok", contextUser))
}

func UserCreate(ctx iris.Context) {
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")
	user, err := models.GetUserByEmail(email)
	if user != nil {
		ctx.JSON(ResponseResource(400, "another email", nil))
	} else {
		user, err = models.CreateUser(email, password)
		if err != nil {
			ctx.JSON(ResponseResource(400, err.Error(), nil))
		} else {
			// 创建完自动登录
			UserLogin(ctx)
		}
	}
}

func UserUpdate(ctx iris.Context) {
	id := ctx.FormValue("id")
	password := ctx.FormValue("password")
	contextUser, err := GetContextUser(ctx)
	if err != nil {
		ctx.JSON(ResponseResource(401, err.Error(), nil))
		return
	}
	if len(contextUser.ID) > 0 {
		id = contextUser.ID
	}
	err = models.UpdatePassword(id, password)
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		// 修改完自动注销
		UserLogout(ctx)
	}
}

func UserPasswordUpdate(ctx iris.Context) {
	id := ctx.FormValue("id")
	password := ctx.FormValue("password")
	contextUser, err := GetContextUser(ctx)
	if err != nil {
		ctx.JSON(ResponseResource(401, err.Error(), nil))
		return
	}
	if len(contextUser.ID) > 0 {
		id = contextUser.ID
	}
	err = models.UpdatePassword(id, password)
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		// 修改完自动注销
		UserLogout(ctx)
	}
}