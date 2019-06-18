package models

import (
	"fmt"
	// "github.com/gin-contrib/sessions"
	// "github.com/gin-contrib/sessions/memstore"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DB *gorm.DB
	// STORE Store
)

func Init() {
	var err error
	DB, err = gorm.Open("sqlite3", "./tmp/test.db")
	if err != nil {
		panic(fmt.Sprintf("No error should happen when connecting to  database, but got err=%+v", err))
	} else {
		fmt.Println("sqlite3 ok")
	}
	DB.LogMode(true)
	DB.AutoMigrate(
		&Role{},
		&UserRole{},
		&User{},
		&BookCatalog{},
		&Book{},
		&Chapter{},
	)
	// store := memstore.NewStore([]byte("secret"))
}

func CloseDB() {
	defer DB.Close()
}
