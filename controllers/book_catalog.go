package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sisimengchen/serverme/models"
	"github.com/sisimengchen/serverme/utils/pagination"
)

func CreateBookCatalog(ctx *gin.Context) {
	name := ctx.PostForm("name")
	if len(name) == 0 {
		ctx.JSON(ResponseResource(400, "require name", nil))
		return
	}
	bookCatalog, err := models.CreateBookCatalog(&models.BookCatalog{Name: name})
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", bookCatalog))
	}
}

func GetBookCatalogByID(ctx *gin.Context) {
	id := ctx.Query("id")
	if len(id) == 0 {
		ctx.JSON(ResponseResource(400, "require id", nil))
		return
	}
	bookCatalog, err := models.GetBookCatalog(&models.BookCatalog{ID: id})
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", bookCatalog))
	}
}

func GetBookCatalogs(ctx *gin.Context) {
	offset, limit := pagination.GetPage(ctx)
	bookCatalogs, err := models.GetBookCatalogs(offset, limit, &models.BookCatalog{})
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", bookCatalogs))
	}
}
