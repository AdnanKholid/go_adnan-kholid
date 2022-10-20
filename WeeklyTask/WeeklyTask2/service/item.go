package service

import (
	"echo-api/models"
	"echo-api/repository"
)

type ItemService struct {
	item repository.ItemRepo
}

func NewItem() ItemService {
	return ItemService{
		item: &repository.ItemRepoImpl{},
	}
}

func (is *ItemService) GetAllItem() []models.Item {
	return is.item.GetAllItem()
}

func (is *ItemService) GetByIDItem(id string) models.Item {
	return is.item.GetByIDItem(id)
}

func (is *ItemService) AddItem(input models.ItemInput) models.Item {
	return is.item.AddItem(input)
}

func (is *ItemService) EditItem(id string, input models.ItemInput) models.Item {
	return is.item.EditItem(id, input)
}

func (is *ItemService) DeleteItem(id string) bool {
	return is.item.DeleteItem(id)
}

func (is *ItemService) GetItemByName(name string) []models.Item {
	return is.item.GetItemByName(name)
}

func (is *ItemService) GetByCategoryID(idCategory string) models.Category {
	return is.item.GetByCategoryID(idCategory)
}
