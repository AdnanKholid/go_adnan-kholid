package controller

import (
	"echo-api/config"
	"echo-api/handler"
	"echo-api/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
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
	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, handler.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	q := config.DB.First(&book, convId).Error
	if q != nil {
		return c.JSON(http.StatusNotFound, handler.Response{
			Message: q.Error(),
			Code:    http.StatusNotFound,
		})
	}

	return c.JSON(http.StatusOK, handler.Response{
		Message: "success",
		Result:  book,
		Code:    http.StatusOK,
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

	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, handler.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	config.DB.Where("id = ?", convId).Delete(&book)

	return c.JSON(http.StatusOK, handler.Response{
		Message: "success",
		Result:  book,
		Code:    http.StatusOK,
	})
}

func UpdateBookController(c echo.Context) error {
	book := model.Book{}
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusUnsupportedMediaType, handler.Response{
			Message: err.Error(),
			Code:    http.StatusUnsupportedMediaType,
		})
	}

	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, handler.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	config.DB.Model(&book).Where("id = ?", convId).Updates(book)

	return c.JSON(http.StatusOK, handler.Response{
		Message: "success",
		Result:  book,
		Code:    http.StatusOK,
	})
}
