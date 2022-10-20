package controller

import (
	"echo-api/models"
	"echo-api/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

var itemController service.ItemService = service.NewItem()

func GetAllItem(c echo.Context) error {
	var items []models.Item = itemController.GetAllItem()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Alert": "Success Get Data",
		"Data":  items,
	})
}

func GetByIDItem(c echo.Context) error {
	var id string = c.Param("id")

	var item models.Item = itemController.GetByIDItem(id)

	if item.ID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Alert": "ID not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Alert": "Success Get Data",
		"Data":  item,
	})
}

func AddItem(c echo.Context) error {
	input := new(models.ItemInput)

	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Alert": "Failed Add Data",
		})
	}
	item := itemController.AddItem(*input)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Alert": "Success Add Data",
		"Data":  item,
	})

}

func EditItem(c echo.Context) error {
	var id string = c.Param("id")

	input := new(models.ItemInput)

	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Alert": "Failed Edit Data",
		})
	}
	item := itemController.EditItem(id, *input)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Alert": "Success Edit Data",
		"Data":  item,
	})

}

func DeleteItem(c echo.Context) error {
	var id string = c.Param("id")

	delItem := itemController.DeleteItem(id)

	if !delItem {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Alert": "Failed Delete Data",
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"Alert": "Success Delete Data",
	})
}

func GetItemByName(c echo.Context) error {
	var name string = c.Param("name")

	item := itemController.GetItemByName(name)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Alert": "Success Get Data",
		"Data":  item,
	})
}

func GetByCategoryID(c echo.Context) error {
	var idCategory string = c.Param("category_id")

	category := itemController.GetByCategoryID(idCategory)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Alert": "Success Get Data",
		"Data":  category,
	})
}
