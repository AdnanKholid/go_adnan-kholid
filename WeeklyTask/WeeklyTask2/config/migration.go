package config

import "echo-api/models"

func AutoMigration() {
	DB.AutoMigrate(&models.Category{}, &models.Item{})
}
