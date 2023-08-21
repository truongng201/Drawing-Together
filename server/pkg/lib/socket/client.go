package socket

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/gommon/log"
)

type Client struct {
	ClientID   	string		`json:"client_id"`
	ClientName 	string			`json:"client_name"`
	Conn 		*websocket.Conn
	WsServer 	*WsServer
	rooms 		map[*Room]bool
	send 		chan []byte
}

func NewClient(Conn *websocket.Conn, clientName string, WsServer *WsServer) *Client {
	return &Client{
		ClientID	:   uuid.New().String(),
		ClientName  : 	clientName,
		Conn		: 	Conn,
		WsServer	: 	WsServer,
		rooms		: 	make(map[*Room]bool),
		send 		: 	make(chan []byte),
	}
}

func (client *Client) disconnect() {
	client.WsServer.Unregister <- client
	for room := range client.rooms {
		room.UnRegister <- client
	}
	close(client.send)
	client.Conn.Close()
}


func (client *Client) ReadMessage() {
	defer func() {
		client.disconnect()
	}()
	for {
		_, message, err := client.Conn.ReadMessage()
		log.Info(message)
		if err != nil {
			log.Error(err)
			return
		}
		client.handleNewMessage(message)
	}
}


func (client *Client) handleNewMessage(jsonMessage []byte) {
	var message Message
	log.Info(jsonMessage)
	if err := json.Unmarshal(jsonMessage, &message); err != nil {
		log.Error(err)
		return
	}

	if message.Sender.ClientName == "" {
		log.Error("Client name is empty")
		return
	}
	client.ClientName = message.Sender.ClientName

	switch message.Action {
	// case ChatAction:
	// 	roomID := message.Target.RoomID
	// 	if room := client.WsServer.FindRoomByID(roomID); room != nil {
	// 		room.Broadcast <- message
	// 	}
	case JoinRoomAction:
		client.handleJoinRoomMessage(message)
	case CreateRoomAction:
		client.handleCreateRoomMessage(message)
	}
}


func (client *Client) WriteMessage() {
	defer func() {
		client.disconnect()
	}()
	
	for {
		select {
		case message := <-client.send:
			log.Info( message)
			var msg Message
			if err := json.Unmarshal(message, &msg); err != nil{
				log.Error(err)
				return  
			}
			if err := client.Conn.WriteJSON(msg); err != nil{
				log.Error(err)
				return
			}
		}
	}
}


func(client *Client) handleJoinRoomMessage(message Message) {
	roomID := message.Target.RoomID
	log.Info(roomID)
	room := client.WsServer.FindRoomByID(roomID);
	log.Info(room)
	if room == nil {
		log.Error("room not found")
		return
	}
	client.rooms[room] = true
	log.Info(client.ClientID)
	room.Register <- client
	room.Broadcast <- Message{
		Action: JoinRoomAction,
		Target: MessageRoom{
			RoomID: room.RoomID,
		},
		Sender: MessageClient{
			ClientName: client.ClientName,
		},
		Payload: "Room joined",
	}
	
}


func(client *Client) handleCreateRoomMessage(message Message) {
	room := NewRoom(message.Target.Private, message.Target.MaxPlayers)
	client.WsServer.Rooms[room] = true
	go room.Start()
	client.rooms[room] = true
	room.Register <- client
	room.Broadcast <- Message{
		Action: CreateRoomAction,
		Target: MessageRoom{
			RoomID: room.RoomID,
			MaxPlayers: room.MaxPlayers,
			Private: room.Private,
		},
		Sender: MessageClient{
			ClientName: client.ClientName,
			ClientID: client.ClientID,
		},
		Payload: "Room created",
	}
}