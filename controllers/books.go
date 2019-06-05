package controllers

import (
	"github.com/kataras/iris"
	"github.com/sisimengchen/serverme/models"
)

func CreateBook(ctx iris.Context) {
	contextUser, err := GetContextUser(ctx)
	if err != nil {
		ctx.JSON(ResponseResource(401, err.Error(), nil))
		return
	}
	name := ctx.FormValue("name")
	if len(name) == 0 {
		ctx.JSON(ResponseResource(400, "require name", nil))
		return
	}
	description := ctx.FormValue("description")
	if len(description) == 0 {
		ctx.JSON(ResponseResource(400, "require description", nil))
		return
	}
	catalogId := ctx.FormValue("catalogId")
	if len(catalogId) == 0 {
		ctx.JSON(ResponseResource(400, "require catalogId", nil))
		return
	}
	authorId := ctx.FormValue("authorId")
	if len(authorId) == 0 {
		authorId = contextUser.ID
	}
	cover := ctx.FormValue("cover")
	book, err := models.CreateBook(&models.Books{Name: name, Description: description, AuthorId: authorId, CatalogId: catalogId, Cover: cover }, contextUser)
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", book))
	}
}

func GetBookByID(ctx iris.Context) {
	_, err := GetContextUser(ctx)
	if err != nil {
		ctx.JSON(ResponseResource(401, err.Error(), nil))
		return
	}
	id := ctx.URLParam("id")
	if len(id) == 0 {
		ctx.JSON(ResponseResource(400, "require id", nil))
		return
	}
	book, err := models.GetBook(&models.Books{ID: id })
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", book))
	}
}

func GetBooksByName(ctx iris.Context) {
	_, err := GetContextUser(ctx)
	if err != nil {
		ctx.JSON(ResponseResource(401, err.Error(), nil))
		return
	}
	name := ctx.URLParam("name")
	if len(name) == 0 {
		ctx.JSON(ResponseResource(400, "require id", nil))
		return
	}
	books, err := models.GetBooks(&models.Books{Name: name })
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", books))
	}
}

func GetBooksByCatalogId(ctx iris.Context) {
	_, err := GetContextUser(ctx)
	if err != nil {
		ctx.JSON(ResponseResource(401, err.Error(), nil))
		return
	}
	catalogId := ctx.URLParam("catalogId")
	if len(catalogId) == 0 {
		ctx.JSON(ResponseResource(400, "require catalogId", nil))
		return
	}
	books, err := models.GetBooks(&models.Books{CatalogId: catalogId })
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", books))
	}
}

func GetBooks(ctx iris.Context) {
	_, err := GetContextUser(ctx)
	if err != nil {
		ctx.JSON(ResponseResource(401, err.Error(), nil))
		return
	}
	books, err := models.GetBooks(&models.Books{})
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", books))
	}
}