package controllers

import (
	"middleware/lib/database"
	"middleware/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetBook(e echo.Context) error {
	books, err := database.GetBook()
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "Books retrieved",
		"data":    books,
	})
}

func GetBookByID(e echo.Context) error {
	id := e.Param("id")

	book, err := database.GetBookByID(id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "Book retrieved",
		"data":    book,
	})
}

func CreateBook(e echo.Context) error {
	book := models.Book{}
	e.Bind(&book)

	newBook, err := database.CreateBook(book)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "Book created",
		"data":    newBook,
	})
}

func DeleteBook(e echo.Context) error {
	id := e.Param("id")

	err := database.DeleteBook(id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "Book deleted",
	})
}

func UpdateBook(e echo.Context) error {
	id := e.Param("id")

	book := models.Book{}
	e.Bind(&book)

	updatedBook, err := database.UpdateBook(id, book)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "Book updated",
		"data":    updatedBook,
	})
}
