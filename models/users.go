package models

import (
	"fmt"
	"time"
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

// 创建用户
func CreateUser(user *Users) (*Users, error) {
	password, err := utils.GeneratePassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = password
	user.ID = utils.GetUUID()
	if err = database.DB.Create(user).Error; err != nil {
		fmt.Printf("CreateUserErr:%s", err)
		return nil, err
	} else {
		return user, nil
	}
}

// 更新用户
func UpdateUser(user *Users) (*Users, error) {
	if len(user.Password) > 0 {
		password, err := utils.GeneratePassword(user.Password)
		if err != nil {
			return nil, err
		}
		user.Password = password
	}
	if err := database.DB.Model(user).Updates(user).Error; err != nil {
		fmt.Printf("UpdateUserErr:%s", err)
		return nil, err
	} else {
		return user, nil
	}
}

// 查询用户
func GetUser(user *Users) (*Users, error) {
	if err := database.DB.Where(user).First(user).Error; err != nil {
		fmt.Printf("GetUserByIdErr:%s", err)
		return nil, err
	}
	return user, nil
}

// 查询所有用户
func GetUsers(user *Users) (*[]Users, error) {
	users := []Users{}
	if err := database.DB.Where(user).Find(&users).Error; err != nil {
		fmt.Printf("GetUserByIdErr:%s", err)
		return nil, err
	}
	return &users, nil
}