package models

import (
	"fmt"
	"time"
	"errors"
	"github.com/sisimengchen/serverme/database"
	"github.com/sisimengchen/serverme/utils"
)

type Users struct {
	ID        string     `gorm:"type:varchar(100);primary_key;" json:"id"`
	Email     string     `gorm:"type:varchar(100);not null;unique_index;" json:"email"`
	Name      string     `gorm:"type:varchar(100)" json:"name"`
	Phone     string     `gorm:"type:varchar(100)" json:"phone"`
	Avatar    string     `gorm:"type:text;" json:"avatar"`
	Password  string     `gorm:"type:varchar(100);not null" json:"-"`
	Role      string     `gorm:"type:varchar(100)" json:"role"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

// id查询用户
func GetUserByID(id string) (*Users, error) {
	if len(id) == 0 {
		return nil, errors.New("require id")
	}
	user := Users{ID: id}
	if err := database.DB.Where(&user).First(&user).Error; err != nil {
		fmt.Printf("GetUserByIdErr:%s", err)
		return nil, err
	}
	return &user, nil
}

// email查询用户
func GetUserByEmail(email string) (*Users, error) {
	if len(email) == 0 {
		return nil, errors.New("require email")
	}
	user := Users{Email: email}
	if err := database.DB.Where(&user).First(&user).Error; err != nil {
		fmt.Printf("GetUserByEmailErr:%s", err)
		return nil, err
	}
	return &user, nil
}

// email+password登录系统
func LoginByEmail(email, password string) (*Users, error) {
	if len(email) == 0 {
		return nil, errors.New("require email")
	}
	if len(password) == 0 {
		return nil, errors.New("require password")
	}
	user := Users{Email: email}
	if err := database.DB.Where(&user).First(&user).Error; err != nil {
		fmt.Printf("LoginByEmailErr:%s", err)
		return nil, err
	}
	isPass, err := utils.ValidatePassword(password, user.Password)
	if isPass {
		fmt.Printf("user:%s", user)
		return &user, nil
	} else {
		fmt.Printf("ValidatePasswordErr:%s", err)
		return nil, err
	}
}

// user模型创建用户表
func CreateUser(email, password string) (*Users, error) {
	if len(email) == 0 {
		return nil, errors.New("require email")
	}
	if len(password) == 0 {
		return nil, errors.New("require password")
	}
	password, err := utils.GeneratePassword(password)
	if (err != nil) {
		return nil, err
	} else {
		user := Users{ID: utils.GetUUID(), Email: email, Password: password}
		if err = database.DB.Create(&user).Error; err != nil {
			fmt.Printf("CreateUserErr:%s", err)
			return nil, err
		} else {
			fmt.Printf("CreateUser:%s", user)
			return &user, nil
		}
	}
}

// 重新设置用户密码
func UpdatePassword(id, password string) error {
	if len(id) == 0 {
		return errors.New("require id")
	}
	if len(password) == 0 {
		return errors.New("require password")
	}
	password, err := utils.GeneratePassword(password)
	if (err != nil) {
		return err
	} else {
		user := Users{ID: id}
		if err = database.DB.Model(&user).Updates(Users{Password: password}).Error; err != nil {
			fmt.Printf("ResetUserPasswordErr:%s", err)
			return err
		} else {
			return nil
		}
	}
}