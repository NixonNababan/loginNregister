package controller

import (
	"loginnregister/models/payload"
	"loginnregister/usecase"
	"net/http"

	"github.com/labstack/echo"
)

func GetCategoryController(c echo.Context) error {
	category, err := usecase.GetCategory()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all category",
		"data":    category,
	})
}

func CreateCategoryController(c echo.Context) error {
	var req payload.CreateCategory
	c.Bind(&req)
	if err := c.Validate(req); err != nil {
		return err
	}
	err := usecase.CreateCategory(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create category",
	})
}
