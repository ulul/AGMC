package controllers

import (
	"middleware/lib/database"
	"middleware/middlewares"
	"middleware/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(e echo.Context) error {
	user := models.User{}
	e.Bind(&user)

	if err := e.Validate(user); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	userLogin, err := database.Login(user)
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
