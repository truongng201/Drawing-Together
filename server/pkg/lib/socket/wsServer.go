package socket

type WsServer struct {
	clients   	map[*Client]bool
	register   	chan *Client
	unregister 	chan *Client
	broadcast  	chan []byte
	rooms      	map[*Room]bool
}

func NewWsServer() *WsServer {
	return &WsServer{
		clients		:   make(map[*Client]bool),
		register	:   make(chan *Client),
		unregister	: 	make(chan *Client),
		broadcast	:  	make(chan []byte),
		rooms		:   make(map[*Room]bool),
	}
}

func (wsServer *WsServer) Start() {
	for {
		select {
		case client := <-wsServer.register:
			wsServer.registerClient(client)
		case client := <-wsServer.unregister:
			wsServer.unregisterClient(client)
		case message := <-wsServer.broadcast:
			wsServer.broadcastToClients(message)
		}
	}
}

func (wsServer *WsServer) registerClient(client *Client)  {
	wsServer.clients[client] = true
}

func (wsServer *WsServer) unregisterClient(client *Client) {
	if _, ok := wsServer.clients[client]; ok {
		delete(wsServer.clients, client)
	}
}

func (wsServer *WsServer) broadcastToClients(message []byte) {
	for client := range wsServer.clients {
		client.send <- message
	}
}

func (wsServer *WsServer) FindRoomByID(roomID string) *Room {
	var foundRoom *Room
	for room := range wsServer.rooms {
		if room.GetRoomID() == roomID {
			foundRoom = room
			break
		}
	}
	return foundRoom	
}


func (wsServer *WsServer) CeateRoom(private bool, maxPlayers int) *Room {
	room := NewRoom(private, maxPlayers)
	go room.Start()
	wsServer.rooms[room] = true
	return room
}