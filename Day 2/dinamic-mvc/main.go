package main

import (
	"mvc/config"
	"mvc/routes"
)

func main() {
	config.InitDB()
	config.InitMigrate()

	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
