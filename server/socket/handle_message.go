package socket

import (
	"cli-typeracer/server/communication"
	"cli-typeracer/server/state"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func handleCommand(ws *websocket.Conn, message *communication.Message) {
	var response communication.Message
	var handleError error

	switch message.Command {
	case communication.Welcome:
		state.CreatePlayer(message.PlayerId, ws)
		response = communication.NewMessage(communication.Welcome, message.PlayerId, "Welcome to the server!")
	case communication.CreateGame:
		var gameId string
		gameId, handleError = state.CreateGame(message.PlayerId)
		response = communication.NewMessage(communication.GameCreated, message.PlayerId, gameId)
	case communication.JoinGame:
		handleError = state.JoinGame(message.Content, message.PlayerId)
		if handleError != nil {
			response = communication.NewMessage(communication.Error, message.PlayerId, handleError.Error())
		}
		response = communication.NewMessage(communication.GameJoined, message.PlayerId, "")
	default:
		response = communication.NewMessage(communication.Unrecognized, message.PlayerId, "Unknown command")
	}

	if handleError != nil {
		response = communication.NewMessage(communication.Error, message.PlayerId, handleError.Error())
		return
	}
	ws.WriteMessage(websocket.TextMessage, []byte(communication.MessageToBytes(&response)))

}

func HandleMessage(ws *websocket.Conn, rawMessage []byte) {
	fmt.Printf("Received message: %s\n", rawMessage)
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
