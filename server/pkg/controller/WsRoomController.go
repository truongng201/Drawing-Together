package controller

import (
	"server/pkg/lib/socket"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type WsRoomController struct {}




func (controller WsRoomController) Execute(c echo.Context, wsServer *socket.WsServer) error {
    conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
    if err != nil {
        return err
    }

    defer conn.Close()

    log.Info("Client address: ",conn.RemoteAddr())


	client := socket.NewClient(
		conn, "", "", wsServer,
	)

	go client.WriteMessage()
	client.ReadMessage()

	wsServer.Register<- client
	return nil
}
