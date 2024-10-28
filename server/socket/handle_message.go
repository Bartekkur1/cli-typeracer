package socket

import (
	"cli-typeracer/server/communication"
	"cli-typeracer/server/handler"
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

func handleCommand(ws *websocket.Conn, message *communication.Message) {
	var response communication.Message
	var handlerError error

	var handler = handler.CommandHandlers[message.Command]
	if handler == nil {
		log.Printf("No handler found for command: %v", message.Command)
		response = communication.NewMessage(communication.Unrecognized, message.PlayerId, "Unknown command")
		ws.WriteMessage(websocket.TextMessage, []byte(communication.MessageToBytes(&response)))
		return
	}

	response, handlerError = handler(ws, message)
	if handlerError != nil {
		response = communication.NewMessage(communication.Error, message.PlayerId, handlerError.Error())
	}

	ws.WriteMessage(websocket.TextMessage, []byte(communication.MessageToBytes(&response)))
}

func HandleMessage(ws *websocket.Conn, rawMessage []byte) {
	var message communication.Message
	var response communication.Message

	err := json.Unmarshal(rawMessage, &message)
	if err != nil {
		log.Printf("Failed to unmarshal message: %v", err)
		response = communication.NewMessage(communication.Error, message.PlayerId, "Invalid JSON")
		ws.WriteMessage(websocket.TextMessage, []byte(communication.MessageToBytes(&response)))
		return
	}

	var validationError = ValidateMessage(&message)
	if validationError != nil {
		response = communication.NewMessage(communication.Error, message.PlayerId, validationError.Error())
		ws.WriteMessage(websocket.TextMessage, []byte(communication.MessageToBytes(&response)))
		return
	}

	handleCommand(ws, &message)
}
