package usecase

import (
	"loginnregister/lib/database"
	"loginnregister/models"
	"loginnregister/models/payload"

	"golang.org/x/crypto/bcrypt"
)

func Login(req *payload.LoginForm) (user models.User, err error) {
	u, e := database.GetUserByEmailOrUsername(req.ID)
	if e != nil {
		err = e
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		return
	}

	user = u
	return
}

func Register(req *payload.Register) error {
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := models.User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: string(password),
	}
	if err := database.Register(&user); err != nil {
		return err
	}
	return nil
}
