package socket

import (
	"cli-typeracer/server/communication"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func HandleMessage(ws *websocket.Conn, rawMessage []byte) {
	fmt.Printf("Received message: %s\n", rawMessage)
	var message communication.Message

	err := json.Unmarshal(rawMessage, &message)
	if err != nil {
		log.Printf("Failed to unmarshal message: %v", err)
		ws.WriteMessage(websocket.TextMessage, []byte("Invalid JSON"))
		return
	}
	var response communication.Message

	var validationError = ValidateMessage(&message)
	if validationError != nil {
		response = communication.NewMessage(communication.Error, message.PlayerId, validationError.Error())
	} else {
		switch message.Command {
		case communication.Welcome:
			response = communication.NewMessage(communication.Welcome, message.PlayerId, "Welcome to the server!")
		default:
			response = communication.NewMessage(communication.Unrecognized, message.PlayerId, "Unknown command")
		}
	}

	ws.WriteMessage(websocket.TextMessage, []byte(communication.MessageToBytes(&response)))
}
