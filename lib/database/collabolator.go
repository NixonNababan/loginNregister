package database

import (
	"loginnregister/config"
	"loginnregister/models"
)

func CreateCollabolator(taskId uint, userId uint) error {
	if err := config.DB.Save(&models.Collaborator{
		UserID: userId,
		TaskID: taskId,
	}).Error; err != nil {
		return err
	}
	return nil
}
