package utils

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// 获取一个uuid字符串
func GetUUID() string {
	return uuid.New().String()
}

// 加密一个密码
func GeneratePassword(userPassword string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	return string(pass), err
}

// 验证一个密码
func ValidatePassword(userPassword, hashed string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, err
	}
	return true, nil
}
