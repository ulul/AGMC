package routes

import (
	"middleware/config"
	"middleware/controllers"
	"middleware/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	middlewares.LogMiddleware(e)

	e.POST("/login", controllers.Login)
	e.GET("/users", controllers.GetUser)
	e.GET("/books", controllers.GetBook)
	e.GET("/books/:id", controllers.GetBookByID)

	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(config.JWTSecret)))

	// user routes
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.GetUserByID)
	r.DELETE("/users/:id", controllers.DeleteUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	// book routes
	r.POST("/books", controllers.CreateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.PUT("/books/:id", controllers.UpdateBook)

	return e
}
