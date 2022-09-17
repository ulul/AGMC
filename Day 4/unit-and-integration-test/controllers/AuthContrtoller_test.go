package controllers

import (
	"testing/models"

	"github.com/stretchr/testify/mock"
)

type MockUser struct {
	mock.Mock
}

func (m *MockUser) Login(user models.User) (interface{}, error) {
	args := m.Called(user)
	return args.Get(0).(models.User), args.Error(1)
}

// func TestLogin(t *testing.T) {
// 	t.Run("User can login", func(t *testing.T) {

// 		userLogin := models.User{
// 			Username: "ulul",
// 			Password: "secret",
// 		}
// 		userDBMock := new(MockUser)
// 		userDBMock.On("Login", mock.Anything).Return(userLogin, nil)

// 		req, err := http.NewRequest(http.MethodPost, "/login", nil)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
// 		rr := httptest.NewRecorder()

// 		handler := h.Login(userLogin)

// 		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
// 		// directly and pass in our Request and ResponseRecorder.
// 		handler.ServeHTTP(rr, req)

// 		// Check the status code is what we expect.
// 		assert.Equal(t, http.StatusOK, rr.Code)

// 		// Check the response body is what we expect.
// 		expected := `{"name": "ulul"}`
// 		assert.JSONEq(t, expected, rr.Body.String())
// 	})
// }
