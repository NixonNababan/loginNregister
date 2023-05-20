package usecase

import (
	"loginnregister/lib/database"
	"loginnregister/models"
	"loginnregister/models/payload"
)

func CreateCategory(req *payload.CreateCategory) error {
	category := models.Category{
		Name: req.Name,
	}
	err := database.CreateCategory(&category)
	if err != nil {
		return err
	}
	return nil
}

func GetCategory() (response []string, err error) {
	category, err := database.GetCategory()
	if err != nil {
		return
	}
	for _, value := range category {
		response = append(response, value.Name)
	}
	return
}
