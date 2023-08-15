package controller

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type WSMessageChatController struct {}

// Define our message object
type Message struct {
    Email    string `json:"email"`
    Username string `json:"username"`
    Message  string `json:"message"`
}

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // broadcast channel


func (controller WSMessageChatController) Execute(c echo.Context) error {
    ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
    if err != nil {
        return err
    }

    defer ws.Close()

    fmt.Println("Client Message Chat connected:", ws.RemoteAddr())

    clients[ws] = true

    for {
		var msg Message

        err := ws.ReadJSON(msg)

        if err != nil {
            c.Logger().Error(err)
            // delete(clients, ws)
			return err
        }
		fmt.Println("Message Received:", msg)

    }
}
