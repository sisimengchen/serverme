package middleware

import (
	// "fmt"
	"github.com/gorilla/securecookie"
	"github.com/kataras/iris"
	"github.com/sisimengchen/serverme/models"
	"github.com/sisimengchen/serverme/utils"
)

var (
	// AES only supports key sizes of 16, 24 or 32 bytes.
	// You either need to provide exactly that amount or you derive the key from what you type in.
	hashKey   = []byte("the-big-and-secret-fash-key-here")
	blockKey  = []byte("lot-secret-of-characters-big-too")
	sc        = securecookie.New(hashKey, blockKey)
	debugUser = models.User{
		Name:     "debugName",
		Username: "debugUsername",
		Password: "debugPassword",
	}
	realUser = models.User{
		Name:     "mengchen",
		Username: "mengchen",
		Password: "mengchen",
	}
)

func Auth(ctx iris.Context) {
	// fmt.Println("Auth middleware")
	// ctx.SetCookieKV("fess", "fess", iris.CookieEncode(sc.Encode))
	// 需要做权限验证
	if utils.IsAuthRequest(ctx) {
		// fmt.Println("Auth valadate")
		debugmode := ctx.GetCookie("debugmode")
		// 如果开启了debug模式
		if debugmode == "debugmode" {
			ctx.Values().Set("user", debugUser)
			ctx.Next()
		} else {
			fess := ctx.GetCookie("fess", iris.CookieDecode(sc.Decode))
			if fess == "fess" {
				ctx.Values().Set("user", realUser)
				ctx.Next()
			} else {
				if utils.IsApiRequest(ctx) {
					ctx.JSON(iris.Map{
						"message": "nologin",
					})
				} else {
					ctx.Redirect("/login")
				}
			}
		}
	} else {
		ctx.Next()
	}

}
