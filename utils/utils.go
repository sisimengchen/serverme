package utils

import (
	"regexp"
	"github.com/kataras/iris"
)

var authRegexp = regexp.MustCompile(`.*/auth/.*`)
var apiRegexp = regexp.MustCompile(`^/api/.*`)

func IsApiRequest(ctx iris.Context) bool {
	return apiRegexp.MatchString(ctx.Path())
}

func IsAuthRequest(ctx iris.Context) bool {
	return authRegexp.MatchString(ctx.Path())
}