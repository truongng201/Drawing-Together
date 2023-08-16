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

func (controller WSMessageChatController) HandleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		fmt.Println("Messages recievied",msg)
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				fmt.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func (controller WSMessageChatController) Execute(c echo.Context) error {
    ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
    if err != nil {
        return err
    }

    defer ws.Close()

    fmt.Println("Client Message Chat connected:", ws.RemoteAddr())

	// Register our new client
	clients[ws] = true

	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Printf("error: %v", err)
			delete(clients, ws)
		}
		// Send the newly received message to the broadcast channel
		fmt.Println("Messages recievied 1",msg)
		
		broadcast <- msg
	}
}
