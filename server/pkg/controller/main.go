package controller

type AppController struct {
	HealthCheckController      HealthCheckController
	WsRoomController           WsRoomController
	CreateRoomController       CreateRoomController
	CheckRoomExistedController CheckRoomExistedController
}
