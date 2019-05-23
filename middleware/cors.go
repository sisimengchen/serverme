package middleware

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/context"
)

// https://github.com/iris-contrib/middleware/blob/master/cors/cors.go

func Cros() context.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
	})
}
