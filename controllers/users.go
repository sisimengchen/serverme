package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sisimengchen/serverme/models"
	"github.com/sisimengchen/serverme/utils"
	"mime/multipart"
	"path/filepath"
)

func Login(ctx *gin.Context) {
	email := ctx.PostForm("email")
	if len(email) == 0 {
		ctx.JSON(ResponseResource(400, "require email", nil))
		return
	}
	password := ctx.PostForm("password")
	if len(password) == 0 {
		ctx.JSON(ResponseResource(400, "require password", nil))
		return
	}
	user, err := models.GetUser(&models.Users{Email: email})
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
		return
	}
	isPass, err := utils.ValidatePassword(password, user.Password)
	if isPass {
		ctx.SetCookie("fess", user.ID, 3600, "/", "", false, true)
		// utils.SetSecureCookie(ctx, "fess", user.ID)
		ctx.JSON(ResponseResource(200, "ok", user))
	} else {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	}
}

func Logout(ctx *gin.Context) {
	ctx.SetCookie("fess", "", -1, "/", "", false, true)
	ctx.JSON(ResponseResource(200, "ok", nil))
}

func GetUser(ctx *gin.Context) {
	id := ctx.PostForm("id")
	if len(id) > 0 {
		user, err := models.GetUser(&models.Users{ID: id})
		if err != nil {
			ctx.JSON(ResponseResource(400, err.Error(), nil))
			return
		}
		ctx.JSON(ResponseResource(200, "ok", user))
	} else {
		contextUser, err := GetContextUser(ctx)
		if err != nil {
			ctx.JSON(ResponseResource(401, err.Error(), nil))
			return
		}
		ctx.JSON(ResponseResource(200, "ok", contextUser))
	}
}

func CreateUser(ctx *gin.Context) {
	email := ctx.PostForm("email")
	if len(email) == 0 {
		ctx.JSON(ResponseResource(400, "require email", nil))
		return
	}
	password := ctx.PostForm("password")
	if len(password) == 0 {
		ctx.JSON(ResponseResource(400, "require password", nil))
		return
	}
	user, _ := models.GetUser(&models.Users{Email: email})
	if user != nil {
		ctx.JSON(ResponseResource(400, "another email", user))
		return
	}
	user, err := models.CreateUser(&models.Users{Email: email, Password: password})
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		// 创建完自动登录
		Login(ctx)
	}
}

func UpdateUserPassword(ctx *gin.Context) {
	id := ctx.PostForm("id")
	if len(id) == 0 {
		ctx.JSON(ResponseResource(400, "require id", nil))
		return
	}
	password := ctx.PostForm("password")
	if len(password) == 0 {
		ctx.JSON(ResponseResource(400, "require password", nil))
		return
	}
	user, err := models.GetUser(&models.Users{ID: id})
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
		return
	}
	user.Password = password
	_, err = models.UpdateUser(user)
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		// 注销完自动登出
		Logout(ctx)
	}
}

func UpdateUser(ctx *gin.Context) {
	id := ctx.PostForm("id")
	if len(id) == 0 {
		ctx.JSON(ResponseResource(400, "require id", nil))
		return
	}
	name := ctx.PostForm("name")
	phone := ctx.PostForm("phone")
	user, err := models.GetUser(&models.Users{ID: id})
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
		return
	}
	user.Name = name
	user.Phone = phone
	user, err = models.UpdateUser(user)
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", user))
	}
}

func UpdateAvatar(ctx *gin.Context) {
	id := ctx.PostForm("id")
	if len(id) == 0 {
		ctx.JSON(ResponseResource(400, "require id", nil))
		return
	}
	user, err := models.GetUser(&models.Users{ID: id})
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
		return
	}
	fileResponses, err := FileUpload(ctx, filepath.Join("./public"), beforeAvatarSave)
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
		return
	}
	user.Avatar = fileResponses[0].Path
	user, err = models.UpdateUser(user)
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", user))
	}
}

func beforeAvatarSave(ctx *gin.Context, file *multipart.FileHeader, fileResponse *FileResponse) {
	file.Filename = ctx.PostForm("id")
	fileResponse.Name = file.Filename
	fileResponse.ContentType = file.Header["Content-Type"][0]
}
