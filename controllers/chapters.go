package controllers

import (
	"path/filepath"
	"mime/multipart"
	"github.com/kataras/iris"
	"github.com/sisimengchen/serverme/models"
)

// 创建文章
func CreateChapter(ctx iris.Context) {
	title := ctx.FormValue("title")
	if len(title) == 0 {
		ctx.JSON(ResponseResource(400, "require title", nil))
		return
	}
	bookId := ctx.FormValue("bookId")
	if len(bookId) == 0 {
		ctx.JSON(ResponseResource(400, "require bookId", nil))
		return
	}
	chapter, err := models.CreateChapter(&models.Chapters{ Title: title, BookId: bookId })
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", chapter))
	}
}

func GetChapterByID(ctx iris.Context) {
	id := ctx.URLParam("id")
	if len(id) == 0 {
		ctx.JSON(ResponseResource(400, "require id", nil))
		return
	}
	chapter, err := models.GetChapter(&models.Chapters{ID: id})
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", chapter))
	}
}

func GetChapterByTitle(ctx iris.Context) {
	title := ctx.URLParam("title")
	if len(title) == 0 {
		ctx.JSON(ResponseResource(400, "require title", nil))
		return
	}
	chapter, err := models.GetChapter(&models.Chapters{Title: title})
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", chapter))
	}
}

func GetChaptersByBookId(ctx iris.Context) {
	bookId := ctx.URLParam("bookId")
	if len(bookId) == 0 {
		ctx.JSON(ResponseResource(400, "require bookId", nil))
		return
	}
	chapters, err := models.GetChapters(&models.Chapters{BookId: bookId})
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", chapters))
	}
}

func GetChapterContent(ctx iris.Context) {
	chapterId := ctx.FormValue("chapterId")
	bookId := ctx.FormValue("bookId")
	path := filepath.Join("./bookstore", bookId, chapterId)
	ctx.ServeFile(path, true)
}

func SetChapterPath(ctx iris.Context) {
	chapterId := ctx.FormValue("chapterId")
	if len(chapterId) == 0 {
		ctx.JSON(ResponseResource(400, "require chapterId", nil))
		return
	}
	bookId := ctx.FormValue("bookId")
	if len(bookId) == 0 {
		ctx.JSON(ResponseResource(400, "require bookId", nil))
		return
	}
	chapter, err := models.GetChapter(&models.Chapters{ID: chapterId, BookId: bookId})
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

func beforeChapterSave(ctx iris.Context, file *multipart.FileHeader, fileResponse *FileResponse) {
	file.Filename = ctx.FormValue("chapterId")
	fileResponse.Name = file.Filename
	fileResponse.ContentType = file.Header["Content-Type"][0]
}

