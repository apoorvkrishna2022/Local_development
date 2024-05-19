package models

import (
	"github.com/apoorvkrishna22/golang_gorm_with_akhil_sharma/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(b)
	return b
}

func GetBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("id = ?", id).First(&getBook)
	return &getBook, db
}

func DeleteBookById(id int64) Book {
	var book Book
	db.Where("id = ?", id).Delete(&book)
	return book
}
