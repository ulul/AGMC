package main

import (
	"middleware/config"
	"middleware/routes"
)

func main() {
	config.InitDB()
	config.InitMigrate()

	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
