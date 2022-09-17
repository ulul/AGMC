package repository

import (
	"testing/models"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (repository *UserRepositoryMock) GetUserByID(id string) (interface{}, error) {
	args := repository.Mock.Called(id)
	if args.Get(0) == nil {
		return nil, nil
	}
	user := args.Get(0).(models.User)
	return user, nil
}

func (repository *UserRepositoryMock) GetUser() (interface{}, error) {
	return "tes", nil
}

func (repository *UserRepositoryMock) CreateUser(user models.User) (interface{}, error) {
	return "tes", nil
}

func (repository *UserRepositoryMock) DeleteUser(id string) error {
	return nil
}
func (repository *UserRepositoryMock) UpdateUser(id string, user models.User) (interface{}, error) {
	return "tes", nil
}
func (repository *UserRepositoryMock) Login(user models.User) (interface{}, error) {
	return "tes", nil
}
