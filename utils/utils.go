package utils

import (
	"github.com/kataras/iris"
	"regexp"
)

var authRegexp = regexp.MustCompile(`.*/auth/.*`)
var apiRegexp = regexp.MustCompile(`^/api/.*`)

func IsApiRequest(ctx iris.Context) bool {
	return apiRegexp.MatchString(ctx.Path())
}

func IsAuthRequest(ctx iris.Context) bool {
	return authRegexp.MatchString(ctx.Path())
}
