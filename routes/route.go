package route

import (
	"loginnregister/controller"
	m "loginnregister/middlewares"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	m.LoggerMiddleware(e)

	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/register", controller.RegisterController)
	e.POST("/login", controller.LoginController)
	e.POST("/logout", controller.LogoutController)

	return e
}
