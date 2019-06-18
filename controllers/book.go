package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sisimengchen/serverme/models"
	"github.com/sisimengchen/serverme/utils/pagination"
)

func CreateBook(ctx *gin.Context) {
	contextUser, err := GetContextUser(ctx)
	if err != nil {
		ctx.JSON(ResponseResource(401, err.Error(), nil))
		return
	}
	name := ctx.PostForm("name")
	if len(name) == 0 {
		ctx.JSON(ResponseResource(400, "require name", nil))
		return
	}
	description := ctx.PostForm("description")
	if len(description) == 0 {
		ctx.JSON(ResponseResource(400, "require description", nil))
		return
	}
	catalogId := ctx.PostForm("catalogId")
	if len(catalogId) == 0 {
		ctx.JSON(ResponseResource(400, "require catalogId", nil))
		return
	}
	authorId := ctx.PostForm("authorId")
	if len(authorId) == 0 {
		authorId = contextUser.ID
	}
	cover := ctx.PostForm("cover")
	book, err := models.CreateBook(&models.Book{Name: name, Description: description, AuthorId: authorId, CatalogId: catalogId, Cover: cover}, contextUser)
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", book))
	}
}

func GetBookByID(ctx *gin.Context) {
	id := ctx.Query("id")
	if len(id) == 0 {
		ctx.JSON(ResponseResource(400, "require id", nil))
		return
	}
	book, err := models.GetBook(&models.Book{ID: id})
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", book))
	}
}

func GetBooksByName(ctx *gin.Context) {
	name := ctx.Query("name")
	if len(name) == 0 {
		ctx.JSON(ResponseResource(400, "require id", nil))
		return
	}
	offset, limit := pagination.GetPage(ctx)
	books, err := models.GetBooks(offset, limit, &models.Book{Name: name})
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", books))
	}
}

func GetBooksByCatalogId(ctx *gin.Context) {
	catalogId := ctx.Query("catalogId")
	if len(catalogId) == 0 {
		ctx.JSON(ResponseResource(400, "require catalogId", nil))
		return
	}
	offset, limit := pagination.GetPage(ctx)
	books, err := models.GetBooks(offset, limit, &models.Book{CatalogId: catalogId})
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", books))
	}
}

func GetBooks(ctx *gin.Context) {
	offset, limit := pagination.GetPage(ctx)
	books, err := models.GetBooks(offset, limit, &models.Book{})
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", books))
	}
}
