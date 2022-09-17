package controllers

import (
	"net/http"
	"testing/lib/database"
	"testing/middlewares"
	"testing/models"

	"github.com/labstack/echo/v4"
)

type AuthControllerHandler interface {
	Login(c echo.Context) error
}

type AuthControllerUserRepository struct {
	userRepo database.UserRepository
}

func NewAuthController(repo database.UserRepository) *AuthControllerUserRepository {
	return &AuthControllerUserRepository{repo}
}

func (h *AuthControllerUserRepository) Login(e echo.Context) error {
	user := models.User{}
	e.Bind(&user)

	if err := e.Validate(user); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	userLogin, err := h.userRepo.Login(user)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	token, _ := middlewares.CreateToken(int(user.ID))

	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "Success Login",
		"user":    userLogin,
		"token":   token,
	})
}
