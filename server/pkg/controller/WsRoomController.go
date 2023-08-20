package controller

import (
	"fmt"
	"server/pkg/lib/socket"

	"github.com/labstack/echo/v4"
)

type WsRoomController struct {}


func (controller WsRoomController) Execute(c echo.Context, wsServer *socket.WsServer) error {
    conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
    if err != nil {
        return err
    }

    defer conn.Close()

    fmt.Println("Client Message Chat connected:", conn.RemoteAddr())


	client := socket.NewClient(
		conn, "Test User", wsServer,
	)

	client.ReadMessage()
	client.WriteMessage()

	wsServer.Register<- client
	return nil
}
