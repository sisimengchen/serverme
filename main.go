package main

import (
	// "fmt"
	"github.com/kataras/iris"
	"github.com/sisimengchen/serverme/app"
	"github.com/sisimengchen/serverme/configs"
)

func main() {
	configs.New()
	app := app.AppInit()

	app.Run(iris.Addr(":8080"), iris.WithPostMaxMemory(32<<20))
}
