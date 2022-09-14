package routes

import (
	"mvc/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controllers.GetUser)
	e.GET("/users/:id", controllers.GetUserByID)
	e.POST("/users", controllers.CreateUser)
	e.DELETE("/users/:id", controllers.DeleteUser)
	e.PUT("/users/:id", controllers.UpdateUser)

	return e
}
