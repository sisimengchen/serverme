package models

import (
	"fmt"
	"time"
	"errors"
	"github.com/sisimengchen/serverme/database"
	"github.com/sisimengchen/serverme/utils" 
)

type Configures struct {
	ID        string     `gorm:"type:varchar(100);primary_key;" json:"id"`
	Toptic    string     `gorm:"type:varchar(100);not null;" json:"toptic"`
	Configure string     `gorm:"type:text;not null" json:"configure"`
	CreatedBy string     `gorm:"type:varchar(100);" json:"createdBy"`
	UpdatedBy string     `gorm:"type:varchar(100);" json:"updatedBy"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-"`
}

// user模型创建用户表
func CreateConfigure(configure *Configures, user *Users) (*Configures, error) {
	if len(configure.Toptic) == 0 {
		return nil, errors.New("require toptic")
	}
	if len(configure.Configure) == 0 {
		return nil, errors.New("require configure")
	}
	if len(configure.ID) == 0 { // id 不存在创建
		configure.ID = utils.GetUUID()
		configure.CreatedBy = user.ID
		if err := database.DB.Create(configure).Error; err != nil {
			fmt.Printf("CreateConfigureErr:%s", err)
			return nil, err
		} else {
			return configure, nil
		}
	} else { // id 存在更新
		fmt.Printf("configure.ID:%s", configure.ID)
		configure.UpdatedBy = user.ID
		if err := database.DB.Model(configure).Updates(configure).Error; err != nil {
			fmt.Printf("ResetUserPasswordErr:%s", err)
			return nil, err
		} else {
			return configure, nil
		}
	}
}

// id查询配置
func GetConfigureByID(id string) (*Configures, error) {
	if len(id) == 0 {
		return nil, errors.New("require id")
	}
	configure := Configures{}
	if err := database.DB.Where(&Configures{ID: id}).First(&configure).Error; err != nil {
		fmt.Printf("GetConfigureByIDErr:%s", err)
		return nil, err
	}
	return &configure, nil
}

// toptic查询所有
func GetConfiguresByTopic(toptic string) (*[]Configures, error) {
	if len(toptic) == 0 {
		return nil, errors.New("require toptic")
	}
	configures := []Configures{}
	if err := database.DB.Where(&Configures{Toptic: toptic}).Find(&configures).Error; err != nil {
		fmt.Printf("GetConfiguresByTopicErr:%s", err)
		return nil, err
	}
	fmt.Printf("GetConfigureByIDErr:%s", configures)
	return &configures, nil
}