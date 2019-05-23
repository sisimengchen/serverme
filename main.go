package main

import (
	"github.com/kataras/iris"
	"github.com/sisimengchen/serverme/app"
)

func main() {

	app := app.AppInit()

	app.Run(iris.Addr(":8080"))
}
