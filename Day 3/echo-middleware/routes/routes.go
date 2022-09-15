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

	// r := e.Group("/jwt")
	// r.Use(middleware.JWT([]byte(config.JWTSecret)))

	// user routes
	e.POST("/users", controllers.CreateUser, middleware.JWT([]byte(config.JWTSecret)))
	e.GET("/users/:id", controllers.GetUserByID, middleware.JWT([]byte(config.JWTSecret)))
	e.DELETE("/users/:id", controllers.DeleteUser, middleware.JWT([]byte(config.JWTSecret)))
	e.PUT("/users/:id", controllers.UpdateUser, middleware.JWT([]byte(config.JWTSecret)))
	// book routes
	e.POST("/books", controllers.CreateBook, middleware.JWT([]byte(config.JWTSecret)))
	e.DELETE("/books/:id", controllers.DeleteBook, middleware.JWT([]byte(config.JWTSecret)))
	e.PUT("/books/:id", controllers.UpdateBook, middleware.JWT([]byte(config.JWTSecret)))

	return e
}
