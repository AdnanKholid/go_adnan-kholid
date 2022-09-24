package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// // get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	var Id string = c.Param("id")
	id, _ := strconv.Atoi(Id)

	for _, user := range users {
		if user.Id == id {
			// splice and replace the user
			return c.JSON(http.StatusOK, user)
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"message": "USER not found",
	})
}

// // delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here

	var Id string = c.Param("id")
	id, _ := strconv.Atoi(Id)

	initialLength := len(users)

	for index, user := range users {
		if user.Id == id {
			// splice and replace the users
			users = append(users[:index], users[index+1:]...)
		}
	}

	if initialLength == len(users) {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "failed delete user",
			"code":     "500",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete user",
		"users":    users,
	})

}

// // update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here

	user := User{}
	c.Bind(&user)

	var Id string = c.Param("id")
	id, _ := strconv.Atoi(Id)

	for _, item := range users {
		if item.Id == id {
			// splice and replace the user
			user.Id = id
			users[id-1] = user
			return c.JSON(http.StatusOK, users[id-1])
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"message": "USER not found",
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.GET("/users/:id", GetUserController)
	e.PUT("/users/:id", UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
