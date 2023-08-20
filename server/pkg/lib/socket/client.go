package socket

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	ClientID   	uuid.UUID		`json:"client_id"`
	ClientName 	string			`json:"client_name"`
	Conn 		*websocket.Conn
	WsServer 	*WsServer
	rooms 		map[*Room]bool
	send 		chan []byte
}

func NewClient(Conn *websocket.Conn, clientName string, WsServer *WsServer) *Client {
	return &Client{
		ClientID	:   uuid.New(),
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
		if err != nil {
			fmt.Println("client.ReadMessage.err", err)
			break
		}
		client.handleNewMessage(message)
	}
}


func (client *Client) handleNewMessage(jsonMessage []byte) {
	var message Message
	if err := json.Unmarshal(jsonMessage, &message); err != nil {
		fmt.Println("client.handleNewMessage.err", err)
		return
	}

	message.Sender = client

	switch message.Action {
	case SendMessageAction:
		roomID := message.Target.RoomID.String()
		if room := client.WsServer.FindRoomByID(roomID); room != nil {
			room.Broadcast <- message
		}
	case JoinRoomAction:
		client.handleJoinRoomMessage(message)
	}
}

func (client *Client) WriteMessage() {
	defer func() {
		client.disconnect()
	}()
	
	for {
		select {
		case message, ok := <-client.send:
			if !ok {
				client.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			client.Conn.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func(client *Client) handleJoinRoomMessage(message Message) {
	roomID := message.Target.RoomID.String()
	room := client.WsServer.FindRoomByID(roomID);
	if room == nil {
		room := NewRoom(false, 5)
		client.WsServer.Rooms[room] = true
		
	}
	client.rooms[room] = true
	room.Register <- client
}