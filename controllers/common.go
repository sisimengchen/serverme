package controllers

import (
	"io"
	"os"
	"path/filepath"
	// "net/http"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sisimengchen/serverme/models"
	"mime/multipart"
	"strings"
)

// 标准化json数据返回
func ResponseResource(code int, msg string, data interface{}) (int, interface{}) {
	return 200, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	}
}

// 获取请求完成路径
func GetContextUser(ctx *gin.Context) (*models.User, error) {
	contextUser, exists := ctx.Get("contextUser")
	if exists {
		user := contextUser.(models.User)
		return &user, nil
	}
	return nil, errors.New("unlogin")
}

type FileResponse struct {
	Index       int    `json:"index"`
	Size        int64  `json:"size"`
	ContentType string `json:"contentType"`
	Name        string `json:"mame"`
	Path        string `json:"path"`
	Url         string `json:"url"`
}

// 通用上传
func Upload(ctx *gin.Context) {
	fileResponses, err := FileUpload(ctx, filepath.Join("./uploads"), beforeSave)
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", fileResponses))
	}
}

func FileUpload(ctx *gin.Context, destDirectory string, before ...func(ctx *gin.Context, file *multipart.FileHeader, fileResponse *FileResponse)) (fileResponses []*FileResponse, err error) {
	form, err := ctx.MultipartForm()
	if err != nil {
		return nil, err
	}
	var totalSize int64
	var count int
	files := form.File["files"]
	for _, file := range files {
		fileResponse := FileResponse{}
		for _, b := range before {
			b(ctx, file, &fileResponse)
		}
		err := saveUploadedFile(file, destDirectory, &fileResponse)
		if err != nil {
			return nil, err
		}
		fileResponse.Index = count
		fileResponses = append(fileResponses, &fileResponse)
		totalSize += fileResponse.Size
		count++
	}
	if count == 0 {
		return nil, errors.New("no file")
	} else {
		return fileResponses, nil
	}
}

func saveUploadedFile(fh *multipart.FileHeader, destDirectory string, fileResponse *FileResponse) (err error) {
	src, err := fh.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	err = os.MkdirAll(destDirectory, os.FileMode(0700))
	if err != nil {
		return err
	}
	path := filepath.Join(destDirectory, fh.Filename)
	out, err := os.OpenFile(path,
		os.O_WRONLY|os.O_CREATE, os.FileMode(0666))
	if err != nil {
		return err
	}
	defer out.Close()
	size, err := io.Copy(out, src)
	if err != nil {
		return err
	}
	fileResponse.Size = size
	fileResponse.Path = strings.Join([]string{"/", path}, "")
	fileResponse.Url = strings.Join([]string{"http://localhost:8080/", path}, "")
	return nil
}

func beforeSave(ctx *gin.Context, file *multipart.FileHeader, fileResponse *FileResponse) {
	fileResponse.Name = file.Filename
	fileResponse.ContentType = file.Header["Content-Type"][0]
}
