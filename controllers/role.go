package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sisimengchen/serverme/models"
	// "github.com/sisimengchen/serverme/utils"
)

func CreateRole(ctx *gin.Context) {
	code := ctx.PostForm("code")
	if len(code) == 0 {
		ctx.JSON(ResponseResource(400, "require code", nil))
		return
	}
	name := ctx.PostForm("name")
	if len(name) == 0 {
		ctx.JSON(ResponseResource(400, "require name", nil))
		return
	}
	role, err := models.CreateRole(&models.Role{Name: name, Code: code})
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", role))
	}
}

// func GetUser(ctx *gin.Context) {
// 	id := ctx.PostForm("id")
// 	if len(id) > 0 {
// 		user, err := models.GetUser(&models.Users{ID: id})
// 		if err != nil {
// 			ctx.JSON(ResponseResource(400, err.Error(), nil))
// 			return
// 		}
// 		ctx.JSON(ResponseResource(200, "ok", user))
// 	} else {
// 		contextUser, err := GetContextUser(ctx)
// 		if err != nil {
// 			ctx.JSON(ResponseResource(401, err.Error(), nil))
// 			return
// 		}
// 		ctx.JSON(ResponseResource(200, "ok", contextUser))
// 	}
// }
