package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func LoggerMiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time = ${time_rfc3339}, host = ${host}, method = ${method}, uri = ${uri}, status = ${status}\n",
	}))
}
