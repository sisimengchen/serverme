package controllers

import (
	"io"
	"os"
	"path/filepath"
	"mime/multipart"
	"errors"
	"github.com/kataras/iris"
	"github.com/sisimengchen/serverme/models"
	"strings"
	"net/http"
	// "fmt"
)

type ResponseJson struct {
	Data   interface{} `json:"data"`
	Code   int         `json:"code"`
	Msg    interface{} `json:"msg"`
}

func ResponseResource(code int, msg string, objects interface{}) (json *ResponseJson) {
	json = &ResponseJson{Code: code, Msg: msg, Data: objects}
	return json
}

// 获取请求完成路径
func GetContextUser(ctx iris.Context) (*models.Users, error) {
	contextUser := ctx.Values().Get("contextUser")
	if ( contextUser != nil) {
		user := contextUser.(models.Users)
		return &user, nil
	}
	return nil, errors.New("unlogin")
}

type FileResponse struct {
	Index        int    `json:"index"`
	Size         int64  `json:"size"`
	ContentType  string `json:"contentType"`
	Name         string `json:"mame"`
	Path         string `json:"path"`
	Url          string `json:"url"`
}

// 通用上传
func Upload(ctx iris.Context)  {
	_, err := GetContextUser(ctx)
	if err != nil {
		ctx.JSON(ResponseResource(401, err.Error(), nil))
		return
	}
	fileResponses, err := FileUpload(ctx, filepath.Join("./uploads"), beforeSave)
	if err != nil {
		ctx.JSON(ResponseResource(400, err.Error(), nil))
	} else {
		ctx.JSON(ResponseResource(200, "ok", fileResponses))
	}
}

// 文件上传公共方法
func FileUpload(ctx iris.Context, destDirectory string, before ...func(ctx iris.Context, file *multipart.FileHeader, fileResponse *FileResponse)) (fileResponses []*FileResponse, err error) {
	maxSize := ctx.Application().ConfigurationReadOnly().GetPostMaxMemory()
	request := ctx.Request()
	err = request.ParseMultipartForm(maxSize)
	if err != nil {
		return nil, err
	}
	var totalSize int64
	var count int
	if request.MultipartForm != nil {
		if fhs := request.MultipartForm.File; fhs != nil {  // map[string][]*FileHeader
			for _, files := range fhs { // 遍历key
				for _, file := range files { // 遍历文件
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
			}
		}
		if count == 0 {
			return nil, errors.New("no file")
		} else {
			return fileResponses, nil
		}
	}
	return nil, http.ErrMissingFile
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
	size , err := io.Copy(out, src)
	if err != nil {
		return err
	}
	fileResponse.Size = size
	fileResponse.Path = strings.Join([]string{"/" , path}, "")
	fileResponse.Url = strings.Join([]string{"http://localhost:8080/" , path}, "")
	return nil
}

func beforeSave(ctx iris.Context, file *multipart.FileHeader, fileResponse *FileResponse) {
	fileResponse.Name = file.Filename
	fileResponse.ContentType = file.Header["Content-Type"][0]
}