package controllers

import (
	"github.com/kataras/iris"
	"github.com/sisimengchen/serverme/models"
)

func CreateBookCatalog(ctx iris.Context) {
	name := ctx.FormValue("name")
	if len(name) == 0 {
		ctx.JSON(ResponseResource(400, "require name", nil))
		return
	}
	bookCatalog, err := models.CreateBookCatalog(&models.BookCatalogs{ Name: name })
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", bookCatalog))
	}
}

func GetBookCatalogByID(ctx iris.Context) {
	id := ctx.URLParam("id")
	if len(id) == 0 {
		ctx.JSON(ResponseResource(400, "require id", nil))
		return
	}
	bookCatalog, err := models.GetBookCatalog(&models.BookCatalogs{ID: id })
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", bookCatalog))
	}
}

func GetBookCatalogs(ctx iris.Context) {
	bookCatalogs, err := models.GetBookCatalogs(&models.BookCatalogs{})
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", bookCatalogs))
	}
}