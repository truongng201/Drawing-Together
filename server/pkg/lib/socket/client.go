package socket

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	ClientID   string `json:"client_id"`
	ClientName string `json:"client_name"`
	AvatarUrl  string `json:"avatar_url"`
	Conn       *websocket.Conn
	WsServer   *WsServer
	rooms      map[*Room]bool
	send       chan []byte
}

func NewClient(Conn *websocket.Conn, clientName string, avatarUrl string, WsServer *WsServer) *Client {
	return &Client{
		ClientID:   uuid.New().String(),
		ClientName: clientName,
		AvatarUrl:  avatarUrl,
		Conn:       Conn,
		WsServer:   WsServer,
		rooms:      make(map[*Room]bool),
		send:       make(chan []byte),
	}
}

func (client *Client) disconnect() {
	client.WsServer.Unregister <- client
	for room := range client.rooms {
		room.UnRegister <- client
	}
	log.Info("Client disconnected")
	close(client.send)
	client.Conn.Close()
}

func (client *Client) ReadMessage() {
	defer func() {
		client.disconnect()
	}()
	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Error(err)
			}
			return
		}
		client.handleNewMessage(message)
	}
}

func (client *Client) handleNewMessage(jsonMessage []byte) {
	var message Message
	if err := json.Unmarshal(jsonMessage, &message); err != nil {
		log.Error(err)
		return
	}

	if message.Sender.ClientName == "" {
		log.Error("Client name is empty")
		return
	}
	client.ClientName = message.Sender.ClientName
	client.AvatarUrl = message.Sender.AvatarUrl

	switch message.Action {
	case ChatAction:
		log.Info("Chat action")
		client.handleChatAction(message)
	case JoinRoomAction:
		client.handleJoinRoomAction(message)
	case CreateRoomAction:
		client.handleCreateRoomAction(message)
	}
}

func (client *Client) WriteMessage() {
	defer func() {
		client.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-client.send:
			if !ok {
				return
			}
			var msg Message

			if err := json.Unmarshal(message, &msg); err != nil {
				log.Error(err)
				return
			}
			if err := client.Conn.WriteJSON(msg); err != nil {
				log.Error(err)
				return
			}
		}
	}
}

func (client *Client) handleJoinRoomAction(message Message) {
	roomID := message.Target.RoomID
	room := client.WsServer.FindRoomByID(roomID)
	if room == nil {
		log.Error("Room not found")
		client.WsServer.Unregister <- client
		client.Conn.Close()
		return
	}
	client.rooms[room] = true
	log.Info(client.ClientID)
	room.Register <- client
	room.Broadcast <- Message{
		Action: JoinRoomAction,
		Target: MessageRoom{
			RoomID:     room.RoomID,
			MaxPlayers: room.MaxPlayers,
			Private:    room.Private,
		},
		Sender: MessageClient{
			ClientName: client.ClientName,
			ClientID:   client.ClientID,
			AvatarUrl:  client.AvatarUrl,
		},
		Payload: "A new user has joined the room",
	}

}

func (client *Client) handleCreateRoomAction(message Message) {
	room := client.WsServer.CreateRoom(message.Target.Private, message.Target.MaxPlayers)
	client.rooms[room] = true
	room.Register <- client
	room.Broadcast <- Message{
		Action: CreateRoomAction,
		Target: MessageRoom{
			RoomID:     room.RoomID,
			MaxPlayers: room.MaxPlayers,
			Private:    room.Private,
		},
		Sender: MessageClient{
			ClientName: client.ClientName,
			ClientID:   client.ClientID,
			AvatarUrl:  client.AvatarUrl,
		},
		Payload: "Room created successfully",
	}
}

func (client *Client) handleChatAction(message Message) {
	roomID := message.Target.RoomID
	room := client.WsServer.FindRoomByID(roomID)
	if room == nil {
		log.Error("Room not found")
		return
	}

	room.Broadcast <- Message{
		Action: ChatAction,
		Target: MessageRoom{
			RoomID:     room.RoomID,
			MaxPlayers: room.MaxPlayers,
			Private:    room.Private,
		},
		Sender: MessageClient{
			ClientName: client.ClientName,
			ClientID:   client.ClientID,
			AvatarUrl:  client.AvatarUrl,
		},
		Payload: message.Payload,
	}
}
