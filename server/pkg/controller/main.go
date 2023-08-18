package controller

import (
	"gorm.io/gorm"
)

type AppController struct {
	HealthCheckController 		HealthCheckController
	CreateRoomController 		CreateRoomController
	WSMessageGuessController 	WSMessageGuessController
	WSMessageChatController 	WSMessageChatController
	GetRoomsController 			GetRoomsController
	Database *gorm.DB
}