package models

import (
	"github.com/jinzhu/gorm"
	"github.com/Anshbir18/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name       string `gorm:"type:varchar(100)" json:"name"`
	Author     string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook()*Book {
	db.NewRecord(b)
	db.Create(&b)

	return b
}

func GetBookById (id int64) (*Book, *gorm.DB) {
	var getbook Book
	db = db.Where("ID = ?", id).Find(&getbook)
	return &getbook, db
}

func GetAllBooks() *[]Book {
	var books []Book
	db.Find(&books)
	return &books
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID = ?", id).Delete(&book)
	return book
}



