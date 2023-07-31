package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func Routes(e *echo.Echo) *echo.Echo {
	e.Use(middleware.Logger())

	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "Ok")
	})

	return e
}