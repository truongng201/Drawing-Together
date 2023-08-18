package controller

import (
	"github.com/labstack/echo/v4"
)

type GetRoomsController struct {}

func (controller GetRoomsController) Execute(c echo.Context) error{
	return c.JSON(200, "Hello from GetRoomsController")
}