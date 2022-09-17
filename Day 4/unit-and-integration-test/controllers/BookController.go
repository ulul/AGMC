package controllers

import (
	"net/http"
	"testing/lib/database"
	"testing/models"

	"github.com/labstack/echo/v4"
)

type BookControllerHandler interface {
	GetBook(c echo.Context) error
	GetBookByID(c echo.Context) error
	CreateBook(c echo.Context) error
	DeleteBook(c echo.Context) error
	UpdateBook(c echo.Context) error
}

type BookControllerBookRepository struct {
	bookRepo database.BookRepository
}

func NewBookController(repo database.BookRepository) *BookControllerBookRepository {
	return &BookControllerBookRepository{repo}
}

func (c *BookControllerBookRepository) GetBook(e echo.Context) error {
	books, err := c.bookRepo.GetBook()
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "Books retrieved",
		"data":    books,
	})
}

func (c *BookControllerBookRepository) GetBookByID(e echo.Context) error {
	id := e.Param("id")

	book, err := c.bookRepo.GetBookByID(id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "Book retrieved",
		"data":    book,
	})
}

func (c *BookControllerBookRepository) CreateBook(e echo.Context) error {
	book := models.Book{}
	e.Bind(&book)

	if err := e.Validate(book); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	newBook, err := c.bookRepo.CreateBook(book)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "Book created",
		"data":    newBook,
	})
}

func (c *BookControllerBookRepository) DeleteBook(e echo.Context) error {
	id := e.Param("id")

	err := c.bookRepo.DeleteBook(id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "Book deleted",
	})
}

func (c *BookControllerBookRepository) UpdateBook(e echo.Context) error {
	id := e.Param("id")

	book := models.Book{}
	e.Bind(&book)

	if err := e.Validate(book); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	updatedBook, err := c.bookRepo.UpdateBook(id, book)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "Book updated",
		"data":    updatedBook,
	})
}
