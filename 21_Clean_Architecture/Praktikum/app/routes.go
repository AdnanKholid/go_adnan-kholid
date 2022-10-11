package routes

import (
	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	LoggerMiddleware echo.MiddlewareFunc
	JWTMilddeware    middleware.JWTConfig
	UserController   users.UserController
}

func (cl *ControllerList) Route(e *echo.Echo) {

	e.GET("/user", cl.UserController.GetAll)
	e.POST("/user", cl.UserController.Create)
}
