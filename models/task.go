package models

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	Title         string `json:"title" form:"title"`
	Description   string `json:"description" form:"description"`
	Status        string `json:"stock" form:"stock"`
	CategoryID    uint   `json:"category_id" form:"category_id"`
	Category      Category
	UserID        uint `json:"user_id" form:"user_id"`
	User          User
	Collaborators []Collaborator
}
