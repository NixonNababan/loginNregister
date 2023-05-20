package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name" gorm:"not null"`
	Username string `json:"username" form:"username" gorm:"unique;not null"`
	Role     string `json:"role" form:"role" gorm:"type:enum('user','admin');default:'user'"`
	Email    string `json:"email" form:"email" gorm:"unique;not null"`
	Password string `json:"password" form:"password"`
}

type Token struct {
	Token string `gorm:"-"`
}
