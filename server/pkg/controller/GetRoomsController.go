package controller

import (
	"server/pkg/lib/database"
	"server/pkg/model"

	"github.com/labstack/echo/v4"
)

type GetRoomsController struct {}

type GetRoomsResponse struct {
	Status 		int 			`json:"status"`
	Message 	string 			`json:"message"`
	Data		[]model.Room 	`json:"data"`
}

func (controller GetRoomsController) Execute(c echo.Context) error{
	db := database.DB.Connect()
	var rooms []model.Room
	db.Find(&rooms)

	return c.JSON(200, &GetRoomsResponse{
		Message: "Success",
		Status: 200,
		Data : rooms,
	})
}