package controllers

import (
	"static-mvc/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllUsers(e echo.Context) error {
	users := models.GetAllUsers()
	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "Users retrieved",
		"data":    users,
	})
}

func GetUser(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	user := models.GetUserByID(id)
	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "Users retrieved",
		"data":    user,
	})
}

func CreateUser(e echo.Context) error {
	user := models.User{}
	e.Bind(&user)
	users := models.CreateUser(user)
	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "User created",
		"data":    users,
	})
}

func UpdateUser(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	user := models.User{}
	e.Bind(&user)
	users := models.UpdateUser(id, user)
	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "User updated",
		"data":    users,
	})
}

func DeleteUser(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	users := models.DeleteUser(id)

	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "User deleted",
		"data":    users,
	})
}
