package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sisimengchen/serverme/utils"
	"time"
)

type Book struct {
	ID           string       `gorm:"type:varchar(36);primary_key;" json:"id"`
	AuthorId     string       `json:"-"`
	Author       *User        `gorm:"association_foreignkey:ID" json:"author"`
	Name         string       `gorm:"type:varchar(36);not null;" json:"name"`
	Description  string       `gorm:"type:varchar(100);not null;" json:"description"`
	State        int          `gorm:"type:int(11);default:'0'" json:"state"`
	Read         int          `gorm:"type:int(11);default:'0'" json:"read"`
	Push         int          `gorm:"type:int(11);default:'0'" json:"push"`
	CatalogId    string       `json:"-"`
	Catalog      *BookCatalog `gorm:"association_foreignkey:ID" json:"catalog"`
	WordCount    int          `gorm:"type:int(11);default:'0'" json:"wordCount"`
	ChapterCount int          `gorm:"type:int(11);default:'0'" json:"chapterCount"`
	Cover        string       `gorm:"type:text;" json:"cover"`
	CreatedBy    string       `gorm:"type:varchar(36);" json:"createdBy"`
	UpdatedBy    string       `gorm:"type:varchar(36);" json:"updatedBy"`
	CreatedAt    *time.Time   `json:"createdAt"`
	UpdatedAt    *time.Time   `json:"updatedAt"`
	DeletedAt    *time.Time   `json:"-"`
}

// 创建图书
func CreateBook(book *Book, user *User) (*Book, error) {
	book.ID = utils.GetUUID()
	book.CreatedBy = user.ID
	if err := DB.Create(book).Error; err != nil {
		fmt.Printf("CreateBookErr:%s", err)
		return nil, err
	} else {
		return book, nil
	}
}

// 更新图书
func UpdateBook(book *Book, user *User) (*Book, error) {
	book.UpdatedBy = user.ID
	if err := DB.Model(book).Updates(book).Error; err != nil {
		fmt.Printf("UpdateBookErr:%s", err)
		return nil, err
	} else {
		return book, nil
	}
}

// 查询图书
func GetBook(book *Book) (*Book, error) {
	if err := DB.Preload("Author").Preload("Catalog").Where(book).First(book).Error; err != nil {
		fmt.Printf("GetBookErr:%s", err)
		return nil, err
	}
	return book, nil
}

// 查询图书列表
func GetBooks(offset int, limit int, book *Book) (*[]Book, error) {
	books := []Book{}
	if err := DB.Preload("Author").Preload("Catalog").Where(book).Offset(offset).Limit(limit).Find(&books).Error; err != nil {
		fmt.Printf("GetBooksErr:%s", err)
		return nil, err
	}
	return &books, nil
}

// 查询热门图书（阅读量排序）
func GetHotBooks(offset int, limit int, book *Book) (*[]Book, error) {
	books := []Book{}
	if err := DB.Preload("Author").Preload("Catalog").Where(book).Offset(offset).Limit(limit).Order("read desc").Find(&books).Error; err != nil {
		fmt.Printf("GetHotBooksErr:%s", err)
		return nil, err
	}
	return &books, nil
}

// 增加阅读量
func AddBookRead(id string) (*Book, error) {
	book := &Book{ID: id}
	if err := DB.Model(book).UpdateColumn("read", gorm.Expr("read + ?", 1)).Error; err != nil {
		fmt.Printf("AddBookReadErr:%s", err)
		return nil, err
	} else {
		return book, nil
	}
}
