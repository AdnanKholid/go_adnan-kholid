package controller

import (
	"echo-api/config"
	"echo-api/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetBooksController(c echo.Context) error {
	var books []model.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   books,
	})
}

func GetBookController(c echo.Context) error {
	var book model.Book

	id, _ := strconv.Atoi(c.Param("id"))

	config.DB.First(&book, "id = ?", id)

	if book.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "Book not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book",
		"user":    book,
	})
}

// create new user
func CreateBookController(c echo.Context) error {
	book := model.Book{}
	c.Bind(&book)

	if err := config.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    book,
	})
}

func DeleteBookController(c echo.Context) error {
	var book model.Book
	c.Bind(&book)

	id, _ := strconv.Atoi(c.Param("id"))
	config.DB.First(&book, "id = ?", id)
	config.DB.Delete(&book)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
	})

}

func UpdateBookController(c echo.Context) error {
	book := model.Book{}
	id, _ := strconv.Atoi(c.Param("id"))
	config.DB.First(&book, id)

	if book.ID == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": "invalid id",
		})
	}
	c.Bind(&book)

	config.DB.Save(&book)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
		"book":    book,
	})
}
