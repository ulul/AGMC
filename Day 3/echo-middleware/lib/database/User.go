package database

import (
	"middleware/config"
	"middleware/models"
)

func GetUser() (interface{}, error) {
	var users []models.User

	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func GetUserByID(id string) (interface{}, error) {
	var user models.User
	if e := config.DB.Where("id = ?", id).First(&user).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func CreateUser(user models.User) (interface{}, error) {
	if e := config.DB.Create(&user).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func DeleteUser(id string) error {
	var user models.User
	if e := config.DB.Where("id = ?", id).Delete(&user).Error; e != nil {
		return e
	}
	return nil
}

func UpdateUser(id string, user models.User) (interface{}, error) {
	if e := config.DB.Where("id = ?", id).Updates(&user).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func Login(user models.User) (interface{}, error) {
	if e := config.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&user).Error; e != nil {
		return nil, e
	}
	return user, nil
}
