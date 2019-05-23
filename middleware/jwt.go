package middleware

import (
	"github.com/kataras/iris/context"
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
)

/**
 * 验证 jwt
 * @method JwtHandler
 */
func Jwt() context.Handler {
	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("serverme"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	return jwtHandler.Serve
}