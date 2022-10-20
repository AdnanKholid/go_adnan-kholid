package routes

import (
	"echo-api/controller"

	"github.com/labstack/echo/v4"
)

func AddRoutes(server *echo.Echo) {
	server.GET("/items", controller.GetAllItem)
	server.GET("/item/:id", controller.GetByIDItem)
	server.GET("/item_by_name/:name", controller.GetItemByName)
	server.GET("/item_by_category_id/:category_id", controller.GetByCategoryID)
	server.POST("/item", controller.AddItem)
	server.PUT("/item/:id", controller.EditItem)
	server.DELETE("/item/:id", controller.DeleteItem)
}
