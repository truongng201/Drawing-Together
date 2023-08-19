package controller

import (
	"fmt"
	"server/pkg/lib/database"
	"server/pkg/model"

	"github.com/labstack/echo/v4"
)

type GetRoomsController struct {
	Room 	model.Room
	Rooms 	[]model.Room
	User 	model.User
	Users 	[]model.User
}


func (controller GetRoomsController) Execute(c echo.Context) error{
	db := database.DB.Connect()
	// jsonMap := map[string]interface{}{}
	res := db.Find(&controller.Users)

	fmt.Printf("%+v\n", res)
	return c.JSON(200, &Response{
		Message: "Success",
		Status: 200,
	})
}