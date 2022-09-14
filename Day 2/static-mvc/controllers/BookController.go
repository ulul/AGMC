package controllers

import (
	"static-mvc/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetBooks(e echo.Context) error {
	users := models.GetAllBooks()
	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "Users retrieved",
		"data":    users,
	})
}

func GetBookByID(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	book := models.GetBookByID(id)
	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "Users retrieved",
		"data":    book,
	})
}

func CreateBook(e echo.Context) error {
	book := models.Book{}
	e.Bind(&book)
	books := models.CreateBook(book)
	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "Book created",
		"data":    books,
	})
}

func UpdateBook(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	book := models.Book{}
	e.Bind(&book)
	books := models.UpdateBook(id, book)
	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "Book updated",
		"data":    books,
	})
}

func DeleteBook(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	book := models.DeleteBook(id)

	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "User deleted",
		"data":    book,
	})
}
