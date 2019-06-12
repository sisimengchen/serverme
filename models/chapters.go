package models

import (
	"fmt"
	"github.com/sisimengchen/serverme/database"
	"github.com/sisimengchen/serverme/utils"
	"time"
)

type Chapters struct {
	ID          string     `gorm:"type:varchar(100);primary_key;" json:"id"`
	Title       string     `gorm:"type:varchar(100);not null;" json:"title"`
	Description string     `gorm:"type:varchar(100);" json:"description"`
	BookId      string     `json:"-"`
	Book        *Books     `gorm:"association_foreignkey:ID" json:"book"`
	Path        string     `gorm:"type:text;" json:"path"`
	State       int        `gorm:"type:int(11);default:'0'" json:"state"`
	WordCount   int        `gorm:"type:int(11);default:'0'" json:"wordCount"`
	CreatedAt   *time.Time `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
	DeletedAt   *time.Time `json:"-"`
}

// 创建章节
func CreateChapter(chapter *Chapters) (*Chapters, error) {
	chapter.ID = utils.GetUUID()
	if err := database.DB.Create(chapter).Error; err != nil {
		fmt.Printf("CreateChapterErr:%s", err)
		return nil, err
	} else {
		return chapter, nil
	}
}

// 更新章节
func UpdateChapter(chapter *Chapters) (*Chapters, error) {
	if err := database.DB.Model(chapter).Updates(chapter).Error; err != nil {
		fmt.Printf("UpdatedChapterErr:%s", err)
		return nil, err
	} else {
		return chapter, nil
	}
}

// 查询章节
func GetChapter(chapter *Chapters) (*Chapters, error) {
	if err := database.DB.Preload("Book").Preload("Book.Author").Preload("Book.Catalog").Where(chapter).First(chapter).Error; err != nil {
		fmt.Printf("GetUserByIdErr:%s", err)
		return nil, err
	}
	return chapter, nil
}

// 查询章节列表
func GetChapters(offset int, limit int, chapter *Chapters) (*[]Chapters, error) {
	chapters := []Chapters{}
	if err := database.DB.Preload("Book").Preload("Book.Author").Preload("Book.Catalog").Where(chapter).Offset(offset).Limit(limit).Find(&chapters).Error; err != nil {
		fmt.Printf("GetUserByIdErr:%s", err)
		return nil, err
	}
	return &chapters, nil
}
