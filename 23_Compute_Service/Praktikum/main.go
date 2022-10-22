package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.oi/gorm"
)

type Item struct {
	gorm.model
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var DB *gorm.DB

func main() {

	// connect to the database
	var err error

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charse=utf8mb4",
		"root",
		"root",
		"336",
		"learn_docker",
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalin("error: ", err)
	}

	DB.AutoMigrate(&Item{})

	log.Println("connected to the database")
	app := echo.New()

	app.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "hello",
		})
	})
}
