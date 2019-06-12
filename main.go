package main

import (
	"github.com/sisimengchen/serverme/app"
	"github.com/sisimengchen/serverme/configs"
)

func main() {
	configs.New()
	router := app.EngineRouter()
	router.Run(configs.Viper.GetString("app.addr"))
}
