package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	Viper = viper.New()
)

func Init() {
	Viper.AddConfigPath("./configs")
	Viper.SetConfigName("config.development")
	Viper.SetConfigType("yaml")
	err := Viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// fmt.Printf("CreateBookErr:%s", Viper.GetString("db.type"))
	// fmt.Printf("CreateBookErr:%s", Viper.GetString("db.path"))
}
