package socket

import (
	"fmt"
)

// Pool is a pool of connections
type Pool struct {
	// Registered connections.
	Register    chan *Client
	UnRegister  chan *Client
	Clients     map[*Client]bool
	Broadcast   chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:    make(chan *Client),
		UnRegister:  make(chan *Client),
		Clients:     make(map[*Client]bool),
		Broadcast:   make(chan Message),
	}
}

func (p *Pool) Start() {
	for {
		select {
		case client := <-p.Register:
			p.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(p.Clients))
			for client := range p.Clients {
				fmt.Println(*client)
				fmt.Println("A new user joined the chat")
			}
		case client := <-p.UnRegister:
			delete(p.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(p.Clients))
			for client := range p.Clients {
				fmt.Println(client)
				fmt.Println("A user left the chat")
			}
		case message := <-p.Broadcast:
			fmt.Println("Sending message to all clients in Pool")
			for client := range p.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}