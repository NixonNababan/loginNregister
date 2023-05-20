package usecase

import (
	"loginnregister/lib/database"
	"loginnregister/models"
	"loginnregister/models/payload"
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(id uint) (payload.GetUser, error) {
	user, err := database.GetUserById(id)
	if err != nil {
		return payload.GetUser{}, err
	}
	resp := payload.GetUser{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
	}
	return resp, nil
}

func UpdateProfil(id uint, req *payload.UpdateUser) error {
	user, err := database.GetUserById(id)
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "The password is wrong")
	}
	if req.NewPassword != req.RetypePassword {
		return echo.NewHTTPError(http.StatusBadRequest, "The password is not match")
	}
	password, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	profil := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(password),
	}
	if err := database.UpdateProfil(&profil, user.Username); err != nil {
		return err
	}
	return nil
}

func DeleteUser(id uint) error {
	if err := database.DeleteUser(id); err != nil {
		return err
	}
	return nil
}
