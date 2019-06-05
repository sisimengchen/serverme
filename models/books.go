package models

import (
	"fmt"
	"time"
	// "errors"
	"github.com/sisimengchen/serverme/database"
	"github.com/sisimengchen/serverme/utils" 
)

type Books struct {
	ID              string        `gorm:"type:varchar(100);primary_key;" json:"id"`
	AuthorId        string        `json:"-"`
	Author          *Users        `gorm:"association_foreignkey:ID" json:"author"`
	Name            string        `gorm:"type:varchar(100);not null;" json:"name"`
	Description     string        `gorm:"type:varchar(100);not null;" json:"description"`
	State           int           `gorm:"type:int(11);default:'0'" json:"state"`
	Read            int           `gorm:"type:int(11);default:'0'" json:"read"`
	Push            int           `gorm:"type:int(11);default:'0'" json:"push"`
	CatalogId       string        `json:"-"`
	Catalog         *BookCatalogs `gorm:"association_foreignkey:ID" json:"catalog"`
	WordCount       int           `gorm:"type:int(11);default:'0'" json:"wordCount"`
	ChapterCount    int           `gorm:"type:int(11);default:'0'" json:"chapterCount"`
	Cover           string        `gorm:"type:text;" json:"cover"`
	CreatedBy       string        `gorm:"type:varchar(100);" json:"createdBy"`
	UpdatedBy       string        `gorm:"type:varchar(100);" json:"updatedBy"`
	CreatedAt       *time.Time    `json:"createdAt"`
	UpdatedAt       *time.Time    `json:"updatedAt"`
	DeletedAt       *time.Time    `json:"-"`
}

// 创建图书
func CreateBook(book *Books, user *Users) (*Books, error) {
	book.ID = utils.GetUUID()
	book.CreatedBy = user.ID
	if err := database.DB.Create(book).Error; err != nil {
		fmt.Printf("CreateBookErr:%s", err)
		return nil, err
	} else {
		return book, nil
	}
}

// 更新图书
func UpdateBook(book *Books, user *Users) (*Books, error) {
	book.UpdatedBy = user.ID
	if err := database.DB.Model(book).Updates(book).Error; err != nil {
		fmt.Printf("UpdateBookErr:%s", err)
		return nil, err
	} else {
		return book, nil
	}
}

// 查询图书
func GetBook(book *Books) (*Books, error) {
	if err := database.DB.Preload("Author").Preload("Catalog").Where(book).First(book).Error; err != nil {
		fmt.Printf("GetBookErr:%s", err)
		return nil, err
	}
	return book, nil
}

// 查询图书列表
func GetBooks(book *Books) (*[]Books, error) {
	books := []Books{}
	if err := database.DB.Preload("Author").Preload("Catalog").Where(book).Find(&books).Error; err != nil {
		fmt.Printf("GetBooksByCatalogIdErr:%s", err)
		return nil, err
	}
	return &books, nil
}

// 创建图书
// func CreateBook(book *Books, user *Users) (*Books, error) {
// 	if len(book.Name) == 0 {
// 		return nil, errors.New("require name")
// 	}
// 	if len(book.Description) == 0 {
// 		return nil, errors.New("require description")
// 	}
// 	if len(book.ID) == 0 {
// 		if len(book.AuthorId) == 0 {
// 			return nil, errors.New("require AuthorId")
// 		}
// 		_, err := GetUserByID(book.AuthorId)
// 		if err != nil {
// 			return nil, errors.New("AuthorId error")
// 		}
// 		if len(book.CatalogId) == 0 {
// 			return nil, errors.New("require CatalogId")
// 		}
// 		_, err = GetBookCatalog(&BookCatalogs{ID: book.CatalogId })
// 		if err != nil {
// 			return nil, errors.New("CatalogId error")
// 		}
// 		book.ID = utils.GetUUID()
// 		book.CreatedBy = user.ID
// 		if err := database.DB.Create(book).Error; err != nil {
// 			fmt.Printf("CreateBookErr:%s", err)
// 			return nil, err
// 		} else {
// 			return book, nil
// 		}
// 	} else { // id 存在更新
// 		book.UpdatedBy = user.ID
// 		if err := database.DB.Model(book).Updates(book).Error; err != nil {
// 			fmt.Printf("CreateBookErr:%s", err)
// 			return nil, err
// 		} else {
// 			return book, nil
// 		}
// 	}
// }

// id查询图书
// func GetBookByID(id string) (*Books, error) {
// 	if len(id) == 0 {
// 		return nil, errors.New("require id")
// 	}
// 	book := Books{ID: id}
// 	if err := database.DB.Preload("Author").Preload("Catalog").Where(&book).First(&book).Error; err != nil {
// 		fmt.Printf("GetUserByIdErr:%s", err)
// 		return nil, err
// 	}
// 	return &book, nil
// }

// // name查询所有图书
// func GetBooksByName(name string) (*[]Books, error) {
// 	if len(name) == 0 {
// 		return nil, errors.New("require name")
// 	}
// 	books := []Books{}
// 	if err := database.DB.Preload("Author").Preload("Catalog").Where(&Books{Name: name}).Find(&books).Error; err != nil {
// 		fmt.Printf("GetBooksByNameErr:%s", err)
// 		return nil, err
// 	}
// 	return &books, nil
// }

// type查询所有图书
// func GetBooksByCatalogId(catalogId string) (*[]Books, error) {
// 	if len(catalogId) == 0 {
// 		return nil, errors.New("require catalogId")
// 	}
// 	books := []Books{}
// 	if err := database.DB.Preload("Author").Preload("Catalog").Where(&Books{CatalogId: catalogId}).Find(&books).Error; err != nil {
// 		fmt.Printf("GetBooksByCatalogIdErr:%s", err)
// 		return nil, err
// 	}
// 	return &books, nil
// }