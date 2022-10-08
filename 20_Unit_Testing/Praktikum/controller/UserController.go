package controller

import (
	"net/http"
	"strconv"

	"echo-api/config"
	"echo-api/handler"
	"echo-api/middleware"
	"echo-api/model"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	var users []model.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	var user model.User
	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, handler.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	q := config.DB.First(&user, convId).Error
	if q != nil {
		return c.JSON(http.StatusNotFound, handler.Response{
			Message: q.Error(),
			Code:    http.StatusNotFound,
		})
	}

	return c.JSON(http.StatusOK, handler.Response{
		Message: "success",
		Result:  user,
		Code:    http.StatusOK,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	var user model.User

	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, handler.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	config.DB.Where("id = ?", convId).Delete(&user)

	return c.JSON(http.StatusOK, handler.Response{
		Message: "success",
		Result:  user,
		Code:    http.StatusOK,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
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

	config.DB.Model(&user).Where("id = ?", convId).Updates(user)

	return c.JSON(http.StatusOK, handler.Response{
		Message: "success",
		Result:  user,
		Code:    http.StatusOK,
	})
}
func UserLoginController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, handler.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	token, err := middleware.CreateToken(user.ID, user.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, handler.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	resp := model.UsersLoginResponse{user.ID, user.Name, user.Email, token}

	return c.JSON(http.StatusOK, handler.Response{
		Message: "success login",
		Code:    http.StatusOK,
		Result:  resp,
	})
}
