package mocks

import (
	"io"
	"net/http/httptest"
	"testing/utils"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type HttpMock struct {
	E *echo.Echo
}

func (http *HttpMock) RequestMock(method, path string, body io.Reader) (echo.Context, *httptest.ResponseRecorder) {
	http.E.Validator = &utils.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(method, path, body)
	rec := httptest.NewRecorder()
	c := http.E.NewContext(req, rec)

	return c, rec
}
