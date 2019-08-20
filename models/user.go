package models

import (
	"fmt"
	"github.com/sisimengchen/serverme/utils"
	"time"
)

type User struct {
	ID        string     `gorm:"type:varchar(36);primary_key;unique_index" json:"id"`
	Email     string     `gorm:"type:varchar(100);not null;unique_index;" json:"email"`
	Name      string     `gorm:"type:varchar(36)" json:"name"`
	Phone     string     `gorm:"type:varchar(36)" json:"phone"`
	Avatar    string     `gorm:"type:text;" json:"avatar"`
	Password  string     `gorm:"type:varchar(100);not null" json:"-"`
	UserRoles []UserRole `gorm:"foreignkey:UserID;association_foreignkey:ID;" json:"roles"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
	Token     string     `gorm:"type:varchar(100);index" json:"token"`
	Expire    time.Time `json:"-"`
}

// 创建用户
func CreateUser(user *User) (*User, error) {
	password, err := utils.GeneratePassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = password
	user.ID = utils.GetUUID()
	if err = DB.Create(user).Error; err != nil {
		fmt.Printf("CreateUserErr:%s", err)
		return nil, err
	} else {
		return user, nil
	}
}

// 更新用户
func UpdateUser(user *User) (*User, error) {
	if len(user.Password) > 0 {
		password, err := utils.GeneratePassword(user.Password)
		if err != nil {
			return nil, err
		}
		user.Password = password
	}
	if err := DB.Model(user).Updates(user).Error; err != nil {
		fmt.Printf("UpdateUserErr:%s", err)
		return nil, err
	} else {
		return user, nil
	}
}

// 查询用户
func GetUser(user *User) (*User, error) {
	if err := DB.Where(user).First(user).Error; err != nil {
		fmt.Printf("GetUserErr:%s", err)
		return nil, err
	}
	return user, nil
}

// 查询用户（附带权限）ss
func GetUserWithRoles(user *User) (*User, error) {
	if err := DB.Preload("UserRoles").Preload("UserRoles.Role").Where(user).First(user).Error; err != nil {
		fmt.Printf("GetUserWithRolesErr:%s", err)
		return nil, err
	}
	return user, nil
}

// 查询所有用户
func GetUsers(offset int, limit int, user *User) (*[]User, error) {
	users := []User{}
	if err := DB.Where(user).Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		fmt.Printf("GetUsersErr:%s", err)
		return nil, err
	}
	return &users, nil
}

func RefreshToken(user *User) (*User, error) {
	user.Token = utils.GetUUID()
	user.Expire = time.Now().Add(time.Hour * 24 * 1)
	if err := DB.Model(user).Updates(user).Error; err != nil {
		fmt.Printf("RefreshTokenErr:%s", err)
		return nil, err
	} else {
		return user, nil
	}
}

// UserRole 用户角色关联实体
type UserRole struct {
	ID        uint32     `gorm:"column:id;primary_key;auto_increment;" json:"-"`
	UserID    string     `gorm:"type:varchar(36);not null;" json:"-"`      // 用户内码
	RoleID    string     `gorm:"type:varchar(36);not null;" json:"roleID"` // 角色内码
	Role      *Role      `gorm:"association_foreignkey:ID" json:"role"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

// 根据查询角色
func GetUserRoles(user *User) (*[]UserRole, error) {
	userRole := UserRole{UserID: user.ID}
	userRoles := []UserRole{}
	if err := DB.Preload("Role").Where(&userRole).Find(&userRoles).Error; err != nil {
		fmt.Printf("GetUserRolesErr:%s", err)
		return nil, err
	}
	return &userRoles, nil
}
