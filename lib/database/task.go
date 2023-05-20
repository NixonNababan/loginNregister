package database

import (
	"loginnregister/config"
	"loginnregister/models"
	"loginnregister/models/payload"
)

func GetTasks() (resp []models.Task, err error) {
	if err := config.DB.Preload("Category").Preload("Collaborators.User").Preload("User").Find(&resp).Error; err != nil {
		return resp, err
	}
	return
}

func GetTask(id uint) (resp models.Task, err error) {
	if err := config.DB.Preload("Category").Preload("Collaborators.User").Preload("User").First(&resp, id).Error; err != nil {
		return resp, err
	}
	return
}

func GetTaskByCategoryId(id uint) (resp models.Task, err error) {
	if err := config.DB.Preload("Category").Preload("Collaborators.User").Preload("User").Where("category_id = ?", id).First(&resp).Error; err != nil {
		return resp, err
	}
	return resp, nil
}

func CreateTask(req *payload.CreateTask, id uint) error {
	var task models.Task
	if err := config.DB.Model(&task).Create(&models.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      "uncompleted",
		CategoryID:  req.CategoryID,
		UserID:      id,
	}).Error; err != nil {
		return err
	}
	return nil
}

func UpdateCompletedStatus(id uint) error {
	var task models.Task
	if err := config.DB.Model(&task).Where("id = ?", id).Update(&models.Task{
		Status: "completed",
	}).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUncompletedStatus(id uint) error {
	var task models.Task
	if err := config.DB.Model(&task).Where("id = ?", id).Update(&models.Task{
		Status: "uncompleted",
	}).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTask(req *payload.CreateTask, id uint) error {
	var item models.Task
	if err := config.DB.Model(&item).Where("id = ?", id).Updates(&models.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		CategoryID:  req.CategoryID,
	}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTask(id uint) error {
	var item models.Task
	if err := config.DB.Delete(&item, id).Error; err != nil {
		return err
	}
	return nil
}
