package main

import (
	"middleware/config"
	"middleware/routes"

	validator "github.com/go-playground/validator/v10"
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
