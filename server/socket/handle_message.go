package socket

import (
	"cli-typeracer/server/handler"
	"log"

	"github.com/bartekkur1/cli-typeracer/contract/communication"

	"github.com/gorilla/websocket"
)

func handleCommand(ws *websocket.Conn, message *communication.Message) {
	var response communication.Message
	var handlerError error

	var commandHandler = handler.CommandHandlers[message.Command]
	if commandHandler == nil && message.Command != communication.Welcome {
		log.Printf("No handler found for command: %v", message.Command)
		response = communication.NewMessage(communication.Unrecognized, message.PlayerId, "Unknown command")
		ws.WriteMessage(websocket.TextMessage, []byte(communication.MessageToBytes(&response)))
		return
	}

	if message.Command == communication.Welcome {
		response, _ = handler.HandleWelcome(ws)
	} else {
		response, handlerError = commandHandler(message)
		if handlerError != nil {
			response = communication.NewMessage(communication.Error, message.PlayerId, handlerError.Error())
			log.Printf("Error handling command: %v", handlerError)
		}
	}

	ws.WriteMessage(websocket.TextMessage, []byte(communication.MessageToBytes(&response)))
}

func HandleMessage(ws *websocket.Conn, message communication.Message) {
	var response communication.Message

	var validationError = ValidateMessage(&message)
	if validationError != nil {
		response = communication.NewMessage(communication.Error, message.PlayerId, validationError.Error())
		ws.WriteMessage(websocket.TextMessage, []byte(communication.MessageToBytes(&response)))
		return
	}

	handleCommand(ws, &message)
}
