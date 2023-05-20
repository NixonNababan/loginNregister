package cookie

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func CreateJWTCookies(c echo.Context, token string) {
	cookie := new(http.Cookie)
	cookie.Name = "JWTCookie"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
}
