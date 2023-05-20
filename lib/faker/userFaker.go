package faker

import (
	"loginnregister/models"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func UserFaker(db *gorm.DB) *models.User {
	password, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	return &models.User{
		Name:     "admin",
		Username: "admin",
		Email:    "admin",
		Password: string(password),
		Role:     "admin",
	}
}
