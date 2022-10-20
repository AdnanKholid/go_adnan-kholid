package repository

import "echo-api/models"

type ItemRepo interface {
	GetAllItem() []models.Item
	GetByIDItem(id string) models.Item
	AddItem(input models.ItemInput) models.Item
	EditItem(id string, input models.ItemInput) models.Item
	DeleteItem(id string) bool
	GetItemByName(name string) []models.Item
	GetByCategoryID(idCategory string) models.Category
}
