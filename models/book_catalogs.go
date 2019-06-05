package models

import (
	"fmt"
	"time"
	"github.com/sisimengchen/serverme/database"
	"github.com/sisimengchen/serverme/utils" 
)

type BookCatalogs struct {
	ID              string     `gorm:"type:varchar(100);primary_key;" json:"id"`
	Name            string     `gorm:"type:varchar(100);not null;unique_index;" json:"name"`
	CreatedAt       *time.Time `json:"-"`
	UpdatedAt       *time.Time `json:"-"`
	DeletedAt       *time.Time `json:"-"`
}

// 创建分类
func CreateBookCatalog(bookCatalog *BookCatalogs) (*BookCatalogs, error) {
	bookCatalog.ID = utils.GetUUID()
	if err := database.DB.Create(bookCatalog).Error; err != nil {
		fmt.Printf("CreateBookCatalogErr:%s", err)
		return nil, err
	} else {
		return bookCatalog, nil
	}
}

// 更新分类
func UpdateBookCatalog(bookCatalog *BookCatalogs) (*BookCatalogs, error) {
	if err := database.DB.Model(bookCatalog).Updates(bookCatalog).Error; err != nil {
		fmt.Printf("UpdatedBookCatalogErr:%s", err)
		return nil, err
	} else {
		return bookCatalog, nil
	}
}

// 查询分类
func GetBookCatalog(bookCatalog *BookCatalogs) (*BookCatalogs, error) {
	if err := database.DB.Where(&bookCatalog).First(&bookCatalog).Error; err != nil {
		fmt.Printf("GetBookCatalogErr:%s", err)
		return nil, err
	}
	return bookCatalog, nil
}

// 获取或有图书分类
func GetBookCatalogs(bookCatalog *BookCatalogs) (*[]BookCatalogs, error) {
	bookCatalogs := []BookCatalogs{}
	if err := database.DB.Where(bookCatalog).Find(&bookCatalogs).Error; err != nil {
		fmt.Printf("GetBookCatalogsErr:%s", err)
		return nil, err
	}
	return &bookCatalogs, nil
}