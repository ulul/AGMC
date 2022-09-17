package repository

import "testing/models"

type UserRepository interface {
	GetUser() (interface{}, error)
	GetUserByID(id string) (interface{}, error)
	CreateUser(user models.User) (interface{}, error)
	DeleteUser(id string) error
	UpdateUser(id string, user models.User) (interface{}, error)
	Login(user models.User) (interface{}, error)
}
