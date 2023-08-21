package socket

import (
	"github.com/labstack/gommon/log"

	"github.com/google/uuid"
)

// Room is a pool of connections
type Room struct {
	// Registered connections.
	RoomID	  	string			`json:"room_id"`
	Private		bool			`json:"private"`
	MaxPlayers	int				`json:"max_players"`
	Register    chan *Client
	UnRegister  chan *Client
	Clients     map[*Client]bool
	Broadcast   chan Message
}

func NewRoom(private bool, maxPlayers int) *Room {
	return &Room{
		RoomID: 	 uuid.New().String(),
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
		log.Info("Room is full")
	} else{
		room.Clients[client] = true
		log.Info("Size of Connection Room: ", len(room.Clients))
		if len(room.Clients) > 1 {
			log.Info("A new user has joined the chat")
		}
	}
	
}

func (room *Room) unregisterClient(client *Client) {
	delete(room.Clients, client)
	log.Info("Size of Connection Room: ", len(room.Clients))
	log.Info("A user left the chat")
}

func (room *Room) broadcastMessage(message []byte) {
	log.Info("Broadcasting message to all clients in room")
	log.Info("Size of Connection Room: ", len(room.Clients))
	log.Info("Room ID: ", room.RoomID)
	for client := range room.Clients {
		client.send <- message
	}
}

func (room *Room) GetRoomID() string {
	return room.RoomID
}