package communication

import "encoding/json"

type Message struct {
	PlayerId string  `json:"playerId"`
	Command  Command `json:"command"`
	Content  string  `json:"content"`
}

type Command string

const (
	Error        Command = "ERROR"
	Unrecognized Command = "UNRECOGNIZED"
	Welcome      Command = "WELCOME"
)

func NewMessage(command Command, playerId, content string) Message {
	return Message{
		PlayerId: playerId,
		Command:  command,
		Content:  content,
	}
}

func MessageToBytes(message *Message) []byte {
	bytes, _ := json.Marshal(message)
	return bytes
}
