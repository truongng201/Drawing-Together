package routes

import (
	"server/pkg/controller"
	"server/pkg/lib/socket"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, controller controller.AppController) *echo.Echo {
	e.GET("/health", func(c echo.Context) error {
		return controller.HealthCheckController.Execute(c)
	})

	wsServer := socket.NewWsServer()
	go wsServer.Start()

	e.GET("/room/", func(c echo.Context) error {
		return controller.WsRoomController.Execute(c, wsServer)
	})

	return e
}
