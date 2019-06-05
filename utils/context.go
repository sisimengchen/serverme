package utils

import (
	"strings"
	// "errors"
	"regexp"
	// "net/http"
	"github.com/gorilla/securecookie"
	"github.com/kataras/iris"
)

var (
	authRegexp = regexp.MustCompile(`.*/auth/.*`)
	apiRegexp  = regexp.MustCompile(`^/api/.*`)
	// AES only supports key sizes of 16, 24 or 32 bytes.
	// You either need to provide exactly that amount or you derive the key from what you type in.
	hashKey  = []byte("the-big-and-secret-fash-key-here")
	blockKey = []byte("lot-secret-of-characters-big-too")
	sc       = securecookie.New(hashKey, blockKey)
)

// 判断是否是api调用
func IsApiRequest(ctx iris.Context) bool {
	return apiRegexp.MatchString(ctx.Path())
}

// 判断是否需要校验登录
func IsAuthRequest(ctx iris.Context) bool {
	return authRegexp.MatchString(ctx.Path())
}

// 获取加密cookie
func GetSecureCookie(ctx iris.Context, name string) string {
	value := ctx.GetCookie(name, iris.CookieDecode(sc.Decode))
	return value
}

// 设置加密cookie
func SetSecureCookie(ctx iris.Context, name, value string) {
	ctx.SetCookieKV(name, value, iris.CookieEncode(sc.Encode))
}

// 获取cookie
func GetCookie(ctx iris.Context, name string) string {
	value := ctx.GetCookie(name)
	return value
}

// 设置cookie
func SetCookie(ctx iris.Context, name, value string) {
	ctx.SetCookieKV(name, value)
}

// 删除cookie
func RemoveCookie(ctx iris.Context, name string) {
	ctx.RemoveCookie(name)
}

// 获取请求完成路径
func GetFullPath(ctx iris.Context) string {
	request := ctx.Request()
	scheme  := "http://"
	if request.TLS != nil {
        scheme = "https://"
    }
	return strings.Join([]string{scheme, request.Host, request.RequestURI}, "")
}
