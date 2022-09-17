package database

import (
	"testing/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	GetBook() (interface{}, error)
	GetBookByID(id string) (interface{}, error)
	CreateBook(book models.Book) (interface{}, error)
	DeleteBook(id string) error
	UpdateBook(id string, book models.Book) (interface{}, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *bookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) GetBook() (interface{}, error) {
	var books []models.Book

	if e := r.db.Find(&books).Error; e != nil {
		return nil, e
	}
	return books, nil
}

func (r *bookRepository) GetBookByID(id string) (interface{}, error) {
	var book models.Book
	if e := r.db.Where("id = ?", id).First(&book).Error; e != nil {
		return nil, e
	}
	return book, nil
}

func (r *bookRepository) CreateBook(book models.Book) (interface{}, error) {
	if e := r.db.Create(&book).Error; e != nil {
		return nil, e
	}
	return book, nil
}

func (r *bookRepository) DeleteBook(id string) error {
	var book models.Book
	if e := r.db.Where("id = ?", id).Delete(&book).Error; e != nil {
		return e
	}
	return nil
}

func (r *bookRepository) UpdateBook(id string, book models.Book) (interface{}, error) {
	if e := r.db.Where("id = ?", id).Updates(&book).Error; e != nil {
		return nil, e
	}
	return book, nil
}
