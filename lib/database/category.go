package database

import (
	"loginnregister/config"
	"loginnregister/models"
)

func CreateCategory(req *models.Category) error {
	if err := config.DB.Save(&req).Error; err != nil {
		return err
	}
	return nil
}

func GetCategory() (category []models.Category, err error) {
	if err := config.DB.Find(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}
