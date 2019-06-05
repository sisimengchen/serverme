package controllers

import (
	"github.com/kataras/iris"
	"github.com/sisimengchen/serverme/models"
)

func ConfigureCreate(ctx iris.Context) {
	contextUser, err := GetContextUser(ctx)
	if err != nil {
		ctx.JSON(ResponseResource(401, err.Error(), nil))
		return
	}
	id := ctx.FormValue("id")
	toptic := ctx.FormValue("toptic")
	config := ctx.FormValue("config")
	configure, err := models.CreateConfigure(&models.Configures{ID: id, Toptic: toptic, Configure: config }, contextUser)
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", configure))
	}
}

func ConfigureGet(ctx iris.Context) {
	_, err := GetContextUser(ctx)
	if err != nil {
		ctx.JSON(ResponseResource(401, err.Error(), nil))
		return
	}
	id := ctx.URLParam("id")
	configure, err := models.GetConfigureByID(id)
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", configure))
	}
}

func ConfiguresGetByTopic(ctx iris.Context) {
	_, err := GetContextUser(ctx)
	if err != nil {
		ctx.JSON(ResponseResource(401, err.Error(), nil))
		return
	}
	topic := ctx.URLParam("topic")
	configures, err := models.GetConfiguresByTopic(topic)
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", configures))
	}
}