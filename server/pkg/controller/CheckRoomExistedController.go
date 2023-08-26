package controller

import (
	"net/http"

	"server/pkg/lib/socket"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type CheckRoomExistedController struct{}

type (
	CheckRoomExistedReqBody struct {
		RoomId string `json:"room_id" validate:"required"`
	}
)

func (controller *CheckRoomExistedController) Execute(c echo.Context, wsServer *socket.WsServer) error {
	var reqBody CheckRoomExistedReqBody
	if err := c.Bind(&reqBody); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Invalid request body",
			"data":    "",
		})
	}

	if err := c.Validate(&reqBody); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Invalid request body",
			"data":    "",
		})
	}

	if room := wsServer.FindRoomByID(reqBody.RoomId); room == nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"success": false,
			"message": "Room not found",
			"data":    "",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Room found",
		"data":    "",
	})
}
