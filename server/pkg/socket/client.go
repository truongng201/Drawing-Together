package socket

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

// Define our message object
type Message struct {
    Email    string `json:"email"`
    Username string `json:"username"`
    Message  string `json:"message"`
}

func (c *Client) Read(){
	defer func() {
		c.Pool.UnRegister <- c
		c.Conn.Close()
	}()

	for {
		var msg Message
		err := c.Conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			return
		}

		message := Message{Email: msg.Email, Username: msg.Username, Message: msg.Message}
		c.Pool.Broadcast <- message
		fmt.Printf("Message Received: %+v\n", message)
	}
}