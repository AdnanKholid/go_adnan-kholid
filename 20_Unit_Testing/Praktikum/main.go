package main

import (
	"echo-api/config"
	"echo-api/router"
)

func main() {
	config.InitDB()

	e := router.New()
	e.Start(":8000")
}
