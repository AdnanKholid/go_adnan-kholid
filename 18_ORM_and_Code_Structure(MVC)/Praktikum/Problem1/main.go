package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

type Response struct {
	Message string
	Result  interface{}
	Code    int
}

func InitDB() {
	config := Config{
		DB_Username: "root",
		DB_Password: "root",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "alta_18",
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.DB_Username, config.DB_Password, config.DB_Host, config.DB_Port, config.DB_Name)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString))
	if err != nil {
		panic(err)
	}
}

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func InitialMigration() {
	DB.AutoMigrate(&User{})
}

// get all users
func GetUsersController(c echo.Context) error {
	var users []User

	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	var user User
	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	q := DB.First(&user, convId).Error
	if q != nil {
		return c.JSON(http.StatusNotFound, Response{
			Message: q.Error(),
			Code:    http.StatusNotFound,
		})
	}

	return c.JSON(http.StatusOK, Response{
		Message: "success",
		Result:  user,
		Code:    http.StatusOK,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	if err := DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	var user User

	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	DB.Where("id = ?", convId).Delete(&user)

	return c.JSON(http.StatusOK, Response{
		Message: "success",
		Result:  user,
		Code:    http.StatusOK,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	user := User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusUnsupportedMediaType, Response{
			Message: err.Error(),
			Code:    http.StatusUnsupportedMediaType,
		})
	}

	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	DB.Model(&user).Where("id = ?", convId).Updates(user)

	return c.JSON(http.StatusOK, Response{
		Message: "success",
		Result:  user,
		Code:    http.StatusOK,
	})
}

func main() {
	e := echo.New()

	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	e.Logger.Fatal(e.Start(":8000"))
}
