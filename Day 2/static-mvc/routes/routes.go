package routes

import (
	"static-mvc/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/users", controllers.GetBooks)
	e.GET("/users/:id", controllers.GetBookByID)
	e.POST("/users", controllers.CreateBook)
	e.PUT("/users/:id", controllers.UpdateBook)
	e.DELETE("/users/:id", controllers.DeleteBook)
	return e

}
