package socket

import (
	"fmt"

	"github.com/google/uuid"
)

// Room is a pool of connections
type Room struct {
	// Registered connections.
	RoomID	  	uuid.UUID		`json:"room_id"`
	Private		bool			`json:"private"`
	MaxPlayers	int				`json:"max_players"`
	Register    chan *Client
	UnRegister  chan *Client
	Clients     map[*Client]bool
	Broadcast   chan Message
}

func NewRoom(private bool, maxPlayers int) *Room {
	return &Room{
		RoomID: 	 uuid.New(),
		Private:	 private,
		MaxPlayers:  maxPlayers,
		Register:    make(chan *Client),
		UnRegister:  make(chan *Client),
		Clients:     make(map[*Client]bool),
		Broadcast:   make(chan Message),
	}
}

func (room *Room) Start() {
	for {
		select {
		case client := <-room.Register:
			room.registerClient(client)
		case client := <-room.UnRegister:
			room.unregisterClient(client)
		case message := <-room.Broadcast:	
			room.broadcastMessage(message.encode())
		}
	}
}

func (room *Room) registerClient(client *Client)  {
	if len(room.Clients) >= room.MaxPlayers {
		fmt.Println("Room is full")
	} else{
		room.Clients[client] = true
		fmt.Println("Size of Connection Room: ", len(room.Clients))
		fmt.Println("A new user joined the chat")
	}
	
}

func (room *Room) unregisterClient(client *Client) {
	delete(room.Clients, client)
	fmt.Println("Size of Connection Room: ", len(room.Clients))
	fmt.Println("A user left the chat")
}

func (room *Room) broadcastMessage(message []byte) {
	fmt.Println("Sending message to all clients in Room")
	for client := range room.Clients {
		if err := client.Conn.WriteJSON(message); err != nil {
			fmt.Println(err)
			return
		}
	}
}

func (room *Room) GetRoomID() string {
	return room.RoomID.String()
}