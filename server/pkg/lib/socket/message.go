package socket

const ChatAction = "chat"
const JoinRoomAction = "join-room"
const LeaveRoomAction = "leave-room"
const CreateRoomAction = "create-room"


type Message struct {
	Action  string  `json:"action"`
	Message string  `json:"message"`
	Target  *Room   `json:"target"`
	Sender  *Client `json:"sender"`
}

type MessageRequestClient struct {
	ClientName string `json:"client_name"`
}

type MessageRequestRoom struct {
	RoomID 		string 	`json:"room_id"`
	MaxPlayers 	int 	`json:"max_players"`
	Private 	bool 	`json:"private"`
}


type MessageRequest struct{
	Action  string  				`json:"action"`
	Sender  MessageRequestClient   	`json:"sender"`
	Target  MessageRequestRoom  	`json:"target"`
	Payload string 					`json:"payload"`
}

type MessageResponse struct{
	Action  string  				`json:"action"`
	Sender  MessageRequestClient   	`json:"sender"`
	Target  MessageRequestRoom  	`json:"target"`
	Payload string 					`json:"payload"`
}

// func (message *Message) encode() []byte {
// 	json, err := json.Marshal(message)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	return json
// }