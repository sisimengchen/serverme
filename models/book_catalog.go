package models

import (
	"fmt"
	"github.com/sisimengchen/serverme/utils"
	"time"
)

type BookCatalog struct {
	ID        string     `gorm:"type:varchar(36);primary_key;" json:"id"`
	Name      string     `gorm:"type:varchar(36);not null;unique_index;" json:"name"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

// 创建分类
func CreateBookCatalog(bookCatalog *BookCatalog) (*BookCatalog, error) {
	bookCatalog.ID = utils.GetUUID()
	if err := DB.Create(bookCatalog).Error; err != nil {
		fmt.Printf("CreateBookCatalogErr:%s", err)
		return nil, err
	} else {
		return bookCatalog, nil
	}
}

// 更新分类
func UpdateBookCatalog(bookCatalog *BookCatalog) (*BookCatalog, error) {
	if err := DB.Model(bookCatalog).Updates(bookCatalog).Error; err != nil {
		fmt.Printf("UpdatedBookCatalogErr:%s", err)
		return nil, err
	} else {
		return bookCatalog, nil
	}
}

// 查询分类
func GetBookCatalog(bookCatalog *BookCatalog) (*BookCatalog, error) {
	if err := DB.Where(&bookCatalog).First(&bookCatalog).Error; err != nil {
		fmt.Printf("GetBookCatalogErr:%s", err)
		return nil, err
	}
	return bookCatalog, nil
}

// 获取或有图书分类
func GetBookCatalogs(offset int, limit int, bookCatalog *BookCatalog) (*[]BookCatalog, error) {
	bookCatalogs := []BookCatalog{}
	if err := DB.Where(bookCatalog).Offset(offset).Limit(limit).Find(&bookCatalogs).Error; err != nil {
		fmt.Printf("GetBookCatalogsErr:%s", err)
		return nil, err
	}
	return &bookCatalogs, nil
}
