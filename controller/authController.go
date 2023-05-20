package controller

import (
	"loginnregister/lib/cookie"
	"loginnregister/middlewares"
	"loginnregister/models/payload"
	"loginnregister/usecase"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func LoginController(c echo.Context) error {
	var loginForm payload.LoginForm
	c.Bind(&loginForm)

	if err := c.Validate(loginForm); err != nil {
		return err
	}

	if _, err := c.Cookie("JWTCookie"); err == nil {
		return echo.NewHTTPError(http.StatusForbidden, "Already logged in")
	}
	user, err := usecase.Login(&loginForm)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	token, err := middlewares.CreateToken(user.Username, user.Role, user.ID)
	if err != nil {
		return err
	}
	cookie.CreateJWTCookies(c, token)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login user",
	})
}

func RegisterController(c echo.Context) error {
	var registerForm payload.Register
	c.Bind(&registerForm)

	_, e := c.Cookie("JWTCookie")

	if e == nil {
		return echo.NewHTTPError(http.StatusForbidden, "Already logged in")
	}

	if err := c.Validate(registerForm); err != nil {
		return err
	}

	if registerForm.Password != registerForm.RetypePassword {
		return echo.NewHTTPError(http.StatusBadRequest, "Password not match")
	}

	err := usecase.Register(&registerForm)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to register",
	})
}

func LogoutController(c echo.Context) error {
	cookie, err := c.Cookie("JWTCookie")
	if err != nil {
		return echo.NewHTTPError(http.StatusNotAcceptable, "Not logged in yet")
	}
	cookie.Expires = time.Now().Add(-1 * time.Hour)
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to logout",
	})
}

func Authorization(c echo.Context) (string, uint) {
	cookie, _ := c.Cookie("JWTCookie")
	token, _ := jwt.Parse(cookie.Value, nil)
	claims, _ := token.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	id := uint(claims["user_id"].(float64))
	return username, id
}
