package repository

import (
	"echo-api/config"
	"echo-api/models"
)

type ItemRepoImpl struct{}

func (iri *ItemRepoImpl) GetAllItem() []models.Item {
	var items []models.Item

	config.DB.Find(&items)

	return items
}

func (iri *ItemRepoImpl) GetByIDItem(id string) models.Item {
	var item models.Item

	config.DB.First(&item, "id=?", id)

	return item
}

func (iri *ItemRepoImpl) AddItem(input models.ItemInput) models.Item {
	var item models.Item = models.Item{
		Nama:        input.Nama,
		Description: input.Description,
		Stok:        input.Stok,
		Harga:       input.Harga,
		CategoryID:  input.CategoryID,
	}
	itemAcc := models.Item{}

	result := config.DB.Create(&item)

	result.Last(&itemAcc)

	return itemAcc
}

func (iri *ItemRepoImpl) EditItem(id string, input models.ItemInput) models.Item {
	var item models.Item = iri.GetByIDItem(id)

	item.Nama = input.Nama
	item.Description = input.Description
	item.Stok = input.Stok
	item.Harga = input.Harga
	item.CategoryID = input.CategoryID

	config.DB.Save(&item)

	return item
}

func (iri *ItemRepoImpl) DeleteItem(id string) bool {
	var item models.Item = iri.GetByIDItem(id)

	result := config.DB.Delete(&item)

	if result.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

func (iri *ItemRepoImpl) GetItemByName(name string) []models.Item {
	var items []models.Item

	config.DB.Where("nama=?", name).Find(&items)

	return items
}

func (iri *ItemRepoImpl) GetByCategoryID(idCategory string) models.Category {
	var category models.Category

	config.DB.Model(&category).Preload("Item").Find(&category)

	return category
}
