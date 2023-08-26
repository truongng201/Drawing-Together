package socket

import (
	"encoding/json"
	"fmt"

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
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure, websocket.CloseNoStatusReceived) {
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
		client.handleChatAction(message)
	case JoinRoomAction:
		client.handleJoinRoomAction(message)
	default:
		log.Error("Invalid action")
		client.WsServer.Unregister <- client
		client.Conn.Close()
		return
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
	room := client.WsServer.FindRoomByID(message.Target.RoomID)
	if room == nil {
		log.Error("Room not found")
		client.WsServer.Unregister <- client
		client.Conn.Close()
		return
	}

	client.rooms[room] = true
	
	
	room.Register <- client

	room.GetAllClientInRoom()
	room.Broadcast <- Message{
		Action: JoinRoomAction,
		Target: MessageRoom{
			RoomID:     room.RoomID,
		},
		Sender: MessageClient{
			ClientName: client.ClientName,
			ClientID:   client.ClientID,
			AvatarUrl:  client.AvatarUrl,
		},
		Payload: MessagePayload{
			Message: fmt.Sprintf("%s joined room", client.ClientName),
		},
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
		},
		Sender: MessageClient{
			ClientName: client.ClientName,
		},
		Payload: message.Payload,
	}
}
