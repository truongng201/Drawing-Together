package controller

import (
	"fmt"
	"server/pkg/lib/socket"

	"github.com/labstack/echo/v4"
)

type WSMessageChatController struct {}


func (controller WSMessageChatController) Execute(c echo.Context, pool *socket.Pool) error {
    ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
    if err != nil {
        return err
    }

    defer ws.Close()

    fmt.Println("Client Message Chat connected:", ws.RemoteAddr())

	client := &socket.Client{
		Conn: ws,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
	return nil
}
