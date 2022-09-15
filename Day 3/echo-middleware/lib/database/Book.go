package database

import (
	"middleware/config"
	"middleware/models"
)

func GetBook() (interface{}, error) {
	var books []models.Book

	if e := config.DB.Find(&books).Error; e != nil {
		return nil, e
	}
	return books, nil
}

func GetBookByID(id string) (interface{}, error) {
	var book models.Book
	if e := config.DB.Where("id = ?", id).First(&book).Error; e != nil {
		return nil, e
	}
	return book, nil
}

func CreateBook(book models.Book) (interface{}, error) {
	if e := config.DB.Create(&book).Error; e != nil {
		return nil, e
	}
	return book, nil
}

func DeleteBook(id string) error {
	var book models.Book
	if e := config.DB.Where("id = ?", id).Delete(&book).Error; e != nil {
		return e
	}
	return nil
}

func UpdateBook(id string, book models.Book) (interface{}, error) {
	if e := config.DB.Where("id = ?", id).Updates(&book).Error; e != nil {
		return nil, e
	}
	return book, nil
}
