package database

import (
	"testing/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUser() (interface{}, error) {
	var users []models.User

	if e := r.db.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func (r *userRepository) GetUserByID(id string) (interface{}, error) {
	var user models.User
	if e := r.db.Where("id = ?", id).First(&user).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func (r *userRepository) CreateUser(user models.User) (interface{}, error) {
	if e := r.db.Create(&user).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func (r *userRepository) DeleteUser(id string) error {
	var user models.User
	if e := r.db.Where("id = ?", id).Delete(&user).Error; e != nil {
		return e
	}
	return nil
}

func (r *userRepository) UpdateUser(id string, user models.User) (interface{}, error) {
	if e := r.db.Where("id = ?", id).Updates(&user).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func (r *userRepository) Login(user models.User) (interface{}, error) {
	if e := r.db.Where("username = ? AND password = ?", user.Username, user.Password).First(&user).Error; e != nil {
		return nil, e
	}
	return user, nil
}
