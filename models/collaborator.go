package models

type Collaborator struct {
	UserID uint `json:"user_id" form:"user_id"`
	User   User
	TaskID uint `json:"task_id" form:"task_id"`
}
