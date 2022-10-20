package main

import (
	"echo-api/config"
	"echo-api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config.Connect()
	e := echo.New()

	routes.AddRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))

}
