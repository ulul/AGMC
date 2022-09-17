package routes

import (
	"testing/config"
	"testing/controllers"
	"testing/lib/database"
	"testing/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	middlewares.LogMiddleware(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Welcome to rest api")
	})

	userRepo := database.NewUserRepository(config.DB)
	bookRepo := database.NewBookRepository(config.DB)

	userController := controllers.NewUserController(userRepo)
	authController := controllers.NewAuthController(userRepo)
	bookController := controllers.NewBookController(bookRepo)

	e.POST("/login", authController.Login)
	e.GET("/users", userController.GetUser)
	e.GET("/books", bookController.GetBook)
	e.GET("/books/:id", bookController.GetBookByID)

	// user routes
	e.POST("/users", userController.CreateUser, middleware.JWT([]byte(config.JWTSecret)))
	e.GET("/users/:id", userController.GetUserByID, middleware.JWT([]byte(config.JWTSecret)))
	e.DELETE("/users/:id", userController.DeleteUser, middleware.JWT([]byte(config.JWTSecret)))
	e.PUT("/users/:id", userController.UpdateUser, middleware.JWT([]byte(config.JWTSecret)))
	// book routes
	e.POST("/books", bookController.CreateBook, middleware.JWT([]byte(config.JWTSecret)))
	e.DELETE("/books/:id", bookController.DeleteBook, middleware.JWT([]byte(config.JWTSecret)))
	e.PUT("/books/:id", bookController.UpdateBook, middleware.JWT([]byte(config.JWTSecret)))

	return e
}
