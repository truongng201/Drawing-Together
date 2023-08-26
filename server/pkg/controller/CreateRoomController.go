package controller

import (
	"net/http"

	"server/pkg/lib/socket"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type CreateRoomController struct{}

type (
	CreateRoomReqBody struct {
		MaxPlayers int  `json:"max_players" validate:"required,min=2,max=20"`
		Private    bool `json:"private"`
	}
	CreateRoomResData struct {
		RoomID     string `json:"room_id"`
		MaxPlayers int    `json:"max_players"`
		Private    bool   `json:"private"`
	}
)

func (controller CreateRoomController) Execute(c echo.Context, wsServer *socket.WsServer) error {
	var reqBody CreateRoomReqBody
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
	room := wsServer.CreateRoom(reqBody.Private, reqBody.MaxPlayers)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Room created",
		"data": CreateRoomResData{
			RoomID:     room.RoomID,
			MaxPlayers: room.MaxPlayers,
			Private:    room.Private,
		},
	})
}
