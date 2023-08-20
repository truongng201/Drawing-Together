package socket

// import (
// 	"encoding/json"
// 	"log"
// )

const SendMessageAction = "send-message"
const JoinRoomAction = "join-room"
const LeaveRoomAction = "leave-room"
const CreateRoomAction = "create-room"


type Message struct {
	Action  string  `json:"action"`
	Message string  `json:"message"`
	Target  Room   `json:"target"`
	Sender  Client `json:"sender"`
}

type MessageBodyClient struct {
	ClientName string `json:"client_name"`
}

type MessageBodyRoom struct {
	RoomID string `json:"room_id"`
}


type MessageBody struct {
	Action  string  `json:"action"`
	Message string  `json:"message"`
	Sender  MessageBodyClient   `json:"sender"`
	Target  MessageBodyRoom   `json:"target"`
}

// func (message *Message) encode() []byte {
// 	json, err := json.Marshal(message)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	return json
// }