package routes

import (
	"weakly-task-3/controllers/blogs"
	"weakly-task-3/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	LoggerMiddleware echo.MiddlewareFunc
	JWTMiddleware    middleware.JWTConfig
	AuthController   users.AuthController
	BlogsController  blogs.BlogsController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	users := e.Group("/user")

	users.POST("/register", cl.AuthController.Register)
	users.POST("/login", cl.AuthController.Login)

	auth := e.Group("/user", middleware.JWTWithConfig(cl.JWTMiddleware))

	auth.POST("/logout", cl.AuthController.Logout)

	blogs := e.Group("blogs", middleware.JWTWithConfig(cl.JWTMiddleware))

	blogs.GET("/get_all", cl.BlogsController.GetAll)
	blogs.GET("/get_by_id/:id", cl.BlogsController.GetByID)
	blogs.POST("/create", cl.BlogsController.Create)
	blogs.PUT("/update/:id", cl.BlogsController.Update)
	blogs.DELETE("/delete/:id", cl.BlogsController.Delete)
	blogs.GET("/get_by_title/:title", cl.BlogsController.GetByTitle)
	blogs.GET("/get_by_category_id/:category_id", cl.BlogsController.GetByCategory)

}
