package models

import (
	"fmt"
	"github.com/sisimengchen/serverme/utils"
	"time"
)

type Role struct {
	ID        string     `gorm:"type:varchar(100);primary_key;" json:"id"`
	Code      string     `gorm:"type:varchar(100);not null;unique_index;" json:"code"`
	Name      string     `gorm:"type:varchar(100);not null;unique_index;" json:"name"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

// 创建角色
func CreateRole(role *Role) (*Role, error) {
	role.ID = utils.GetUUID()
	if err := DB.Create(role).Error; err != nil {
		fmt.Printf("CreateRoleErr:%s", err)
		return nil, err
	} else {
		return role, nil
	}
}
