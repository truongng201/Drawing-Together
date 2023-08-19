package controller

type AppController struct {
	HealthCheckController 		HealthCheckController
	CreateRoomController 		CreateRoomController
	WSMessageGuessController 	WSMessageGuessController
	WSMessageChatController 	WSMessageChatController
	GetRoomsController 			GetRoomsController
}