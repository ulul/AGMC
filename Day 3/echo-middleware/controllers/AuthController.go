package controllers

import (
	"middleware/lib/database"
	"middleware/middlewares"
	"middleware/models"

	"github.com/labstack/echo/v4"
)

func Login(e echo.Context) error {
	user := models.User{}
	e.Bind(&user)

	userLogin, err := database.Login(user)
	if err != nil {
		return e.JSON(500, err)
	}

	token, _ := middlewares.CreateToken(int(user.ID))

	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "Success Login",
		"user":    userLogin,
		"token":   token,
	})
}
