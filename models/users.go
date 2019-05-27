package models

import (
	"time"
)

import (
	"fmt"
	// "time"
	// "github.com/jinzhu/gorm"
	"github.com/sisimengchen/serverme/database"
)

type User struct {
	ID        string `gorm:"type:varchar(100);primary_key"`
	Email     string `gorm:"type:varchar(100);not null;unique_index;"`
	Name      string `gorm:"type:varchar(100)"`
	Phone     string `gorm:"type:varchar(100)"`
	Password  string `gorm:"type:varchar(100);not null"`
	Role      string `gorm:"type:varchar(100)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func GetUserById(id string) *User {
	if id == "" {
		return nil
	}
	user := &User{ID: id}
	if err := database.DB.Where(user).First(user).Error; err != nil {
		fmt.Printf("GetUserByIdErr:%s", err)
		return nil
	}
	return user
}

func GetUserByEmail(email string) *User {
	if email == "" {
		return nil
	}
	user := &User{Email: email}
	if err := database.DB.Where(user).First(user).Error; err != nil {
		fmt.Printf("GetUserByEmailErr:%s", err)
		return nil
	}
	return user
}

func LoginByEmail(email, password string) *User {
	if email == "" || password == "" {
		return nil
	}
	user := &User{Email: email, Password: password}
	if err := database.DB.Where(user).First(user).Error; err != nil {
		fmt.Printf("LoginByEmailErr:%s", err)
		return nil
	}
	fmt.Printf("user:%s", user)
	return user
}

func CreateUser(user *User) *User {
	if err := database.DB.Save(user).Error; err != nil {
		fmt.Printf("CreateUserErr:%s", err)
		return nil
	}
	return user
}
