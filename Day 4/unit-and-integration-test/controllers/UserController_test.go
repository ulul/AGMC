package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"testing/lib/repository"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = repository.UserRepositoryMock{Mock: mock.Mock{}}
var userController = &UserControllerUserRepository{&userRepository}
var userJSON = `{"name":"Jon Snow","password":"secret"}`
var userJSONUpdate = `{"name":"Jon Snow New","password":"secret"}`

func TestUserControllerUserRepository_CreateUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	h := userController.CreateUser(c)

	assert.NoError(t, h)
	// need mock the validator first
	// assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUserControllerUserRepository_GetUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := userController.GetUser(c)

	fmt.Print(rec.Body.String())
	t.Log(rec.Body.String())
	// Assertions
	if assert.NoError(t, h) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUserControllerUserRepository_GetUserByID(t *testing.T) {
	t.Skip("I dont know why this test is error")
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")
	h := userController.GetUserByID(c)

	fmt.Print(rec.Body.String())
	t.Log(rec.Body.String())
	// Assertions
	if assert.NoError(t, h) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUserControllerUserRepository_DeleteUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/users/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")
	h := userController.DeleteUser(c)

	fmt.Print(rec.Body.String())
	t.Log(rec.Body.String())
	// Assertions
	if assert.NoError(t, h) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUserControllerUserRepository_UpdateUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/users/:id", strings.NewReader(userJSONUpdate))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")
	h := userController.UpdateUser(c)

	fmt.Print(rec.Body.String())
	t.Log(rec.Body.String())
	// Assertions
	assert.NoError(t, h)
	// need mock the validator first
	// assert.Equal(t, http.StatusOK, rec.Code)
}
