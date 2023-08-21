package socket

import (
	"encoding/json"

	"github.com/labstack/gommon/log"
)

const ChatAction = "chat"
const JoinRoomAction = "join-room"
const LeaveRoomAction = "leave-room"
const CreateRoomAction = "create-room"


type Message struct {
	Action  	string  		`json:"action"`
	Target  	MessageRoom   	`json:"target"`
	Sender  	MessageClient 	`json:"sender"`
	Payload 	string  		`json:"payload"`
}

type MessageClient struct {
	ClientName string `json:"client_name"`
	ClientID   string `json:"client_id"`
}

type MessageRoom struct {
	RoomID 		string 	`json:"room_id"`
	MaxPlayers 	int 	`json:"max_players"`
	Private 	bool 	`json:"private"`
}



func (message *Message) encode() []byte {
	json, err := json.Marshal(message)
	
	if err != nil {
		log.Error(err)
		return nil
	}

	return json
}