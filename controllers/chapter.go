package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sisimengchen/serverme/models"
	"github.com/sisimengchen/serverme/utils/pagination"
	"mime/multipart"
	"path/filepath"
)

// 创建文章
func CreateChapter(ctx *gin.Context) {
	title := ctx.PostForm("title")
	if len(title) == 0 {
		ctx.JSON(ResponseResource(400, "require title", nil))
		return
	}
	bookId := ctx.PostForm("bookId")
	if len(bookId) == 0 {
		ctx.JSON(ResponseResource(400, "require bookId", nil))
		return
	}
	chapter, err := models.CreateChapter(&models.Chapter{Title: title, BookId: bookId})
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", chapter))
	}
}

func GetChapterByID(ctx *gin.Context) {
	id := ctx.Query("id")
	if len(id) == 0 {
		ctx.JSON(ResponseResource(400, "require id", nil))
		return
	}
	chapter, err := models.GetChapter(&models.Chapter{ID: id})
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", chapter))
	}
}

func GetChapterByTitle(ctx *gin.Context) {
	title := ctx.Query("title")
	if len(title) == 0 {
		ctx.JSON(ResponseResource(400, "require title", nil))
		return
	}
	chapter, err := models.GetChapter(&models.Chapter{Title: title})
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", chapter))
	}
}

func GetChaptersByBookId(ctx *gin.Context) {
	bookId := ctx.Query("bookId")
	if len(bookId) == 0 {
		ctx.JSON(ResponseResource(400, "require bookId", nil))
		return
	}
	offset, limit := pagination.GetPage(ctx)
	chapters, err := models.GetChapters(offset, limit, &models.Chapter{BookId: bookId})
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", chapters))
	}
}

func GetChapterContent(ctx *gin.Context) {
	chapterId := ctx.Query("chapterId")
	bookId := ctx.Query("bookId")
	path := filepath.Join("./bookstore", bookId, chapterId)
	ctx.File(path)
}

func SetChapterPath(ctx *gin.Context) {
	chapterId := ctx.PostForm("chapterId")
	if len(chapterId) == 0 {
		ctx.JSON(ResponseResource(400, "require chapterId", nil))
		return
	}
	bookId := ctx.PostForm("bookId")
	if len(bookId) == 0 {
		ctx.JSON(ResponseResource(400, "require bookId", nil))
		return
	}
	chapter, err := models.GetChapter(&models.Chapter{ID: chapterId, BookId: bookId})
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
		return
	}
	fileResponses, err := FileUpload(ctx, filepath.Join("./bookstore", bookId), beforeChapterSave)
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
		return
	}
	chapter.Path = fileResponses[0].Path
	chapter, err = models.UpdateChapter(chapter)
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", chapter))
	}
}

func beforeChapterSave(ctx *gin.Context, file *multipart.FileHeader, fileResponse *FileResponse) {
	file.Filename = ctx.PostForm("chapterId")
	fileResponse.Name = file.Filename
	fileResponse.ContentType = file.Header["Content-Type"][0]
}
