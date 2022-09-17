package controllers

import (
	"net/http"
	"testing/lib/database"
	"testing/models"

	"github.com/labstack/echo/v4"
)

type UserControllerHandler interface {
	GetUser(c echo.Context) error
	GetUserByID(c echo.Context) error
	CreateUser(c echo.Context) error
	GetUserById(c echo.Context) error
	DeleteUser(c echo.Context) error
	UpdateUser(c echo.Context) error
}

type UserControllerUserRepository struct {
	userRepo database.UserRepository
}

func NewUserController(repo database.UserRepository) *UserControllerUserRepository {
	return &UserControllerUserRepository{repo}
}

func (c *UserControllerUserRepository) GetUser(e echo.Context) error {
	users, err := c.userRepo.GetUser()
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "Users retrieved",
		"data":    users,
	})
}

func (c *UserControllerUserRepository) GetUserByID(e echo.Context) error {
	id := e.Param("id")

	user, err := c.userRepo.GetUserByID(id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "User retrieved",
		"data":    user,
	})
}

func (c *UserControllerUserRepository) CreateUser(e echo.Context) error {
	user := models.User{}
	e.Bind(&user)

	if err := e.Validate(user); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	newUser, err := c.userRepo.CreateUser(user)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "User created",
		"data":    newUser,
	})
}

func (c *UserControllerUserRepository) DeleteUser(e echo.Context) error {
	id := e.Param("id")

	err := c.userRepo.DeleteUser(id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "User deleted",
	})
}

func (c *UserControllerUserRepository) UpdateUser(e echo.Context) error {
	id := e.Param("id")
	user := models.User{}
	e.Bind(&user)

	if err := e.Validate(user); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	updatedUser, err := c.userRepo.UpdateUser(id, user)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "User updated",
		"data":    updatedUser,
	})
}
