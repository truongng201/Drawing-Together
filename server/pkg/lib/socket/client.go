package socket

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	ClientID   	uuid.UUID		`json:"client_id"`
	ClientName 	string			`json:"client_name"`
	Conn 		*websocket.Conn
	wsServer 	*WsServer
	rooms 		map[*Room]bool
	send 		chan []byte
}

func NewClient(conn *websocket.Conn, room *Room, clientName string, wsServer *WsServer) *Client {
	return &Client{
		ClientID	:   uuid.New(),
		ClientName  : 	clientName,
		Conn		: 	conn,
		wsServer	: 	wsServer,
		rooms		: 	make(map[*Room]bool),
		send 		: 	make(chan []byte),
	}
}

func (client *Client) disconnect() {
	client.wsServer.unregister <- client
	for room := range client.rooms {
		room.UnRegister <- client
	}
	close(client.send)
	client.Conn.Close()
}

func (client *Client) read() {
	defer func() {
		client.disconnect()
	}()
	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			client.disconnect()
			break
		}
		client.handleMessage(message)
	}
}

func (client *Client) handleMessage(jsonMessage []byte) {
	var message Message
	if err := json.Marshal(jsonMessage)
}
