package models

import (
	"github.com/Tejsvapandey1/goCRUD.git/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Name string `json:"name" gorm:"not null"`
	Author string `json:"author" gorm:"not null`
	Publication string `json:"publication" gorm:"default:'unknown'`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}	

func (b * Book) CreateBook() *Book{	
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books 
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book
	db := db.Find(&getBook,"ID = ?",Id)
	return &getBook,db
}

func DeleteBook(Id int64) Book{
	var book Book
	db.Where("ID=?",Id).Delete(book)
	return book
}