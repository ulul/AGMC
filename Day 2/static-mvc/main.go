package main

import "static-mvc/routes"

func main() {
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
