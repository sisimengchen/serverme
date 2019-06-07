package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DB = New()
)

/**
*设置数据库连接
*@param diver string
 */
func New() *gorm.DB {
	DB, err := gorm.Open("sqlite3", "./tmp/test.db")
	if err != nil {
		panic(fmt.Sprintf("No error should happen when connecting to  database, but got err=%+v", err))
	} else {
		fmt.Println("sqlite3 ok")
	}
	DB.LogMode(true)
	return DB
}
