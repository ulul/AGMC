package main

import (
	"testing/config"
	"testing/routes"

	validator "github.com/go-playground/validator"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	config.InitDB()
	config.InitMigrate()

	e := routes.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Logger.Fatal(e.Start(":8080"))
}
