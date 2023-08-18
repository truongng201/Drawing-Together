package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CreateRoomController struct {}

type Response struct {
	Status 	int `json:"status"`
	Message string `json:"message"`
}

func (controller *CreateRoomController) Execute(c echo.Context) error {
	jsonMap := make(map[string]interface{})
	
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)

	if err != nil {
		fmt.Println("controller.CreateRoomController.Execute", err)
		return c.JSON(http.StatusBadRequest, &Response{
			Message: string(err.Error()),
			Status: http.StatusBadRequest,
		})
	}

	username := jsonMap["username"]
	if username == nil || len(username.(string)) < 6{
		return c.JSON(http.StatusBadRequest, &Response{
			Message: "Invalid username",
			Status: http.StatusBadRequest,
		})
	}


	return c.JSON(http.StatusOK, &Response{
		Message: "Success",
		Status: http.StatusOK,
	})
}