package controller

type AppController struct {
	HealthCheckController HealthCheckController
	WSMessageGuessController WSMessageGuessController
	WSMessageChatController WSMessageChatController
}