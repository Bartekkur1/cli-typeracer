package handler

import (
	"cli-typeracer/server/communication"
	"cli-typeracer/server/state"

	"github.com/gorilla/websocket"
)

func HandleWelcome(ws *websocket.Conn, message *communication.Message) (communication.Message, error) {
	var err = state.CreatePlayer(message.PlayerId, ws)

	if err != nil {
		return communication.Message{}, err
	}

	return communication.NewMessage(communication.Welcome, message.PlayerId, "Welcome to the server!"), nil
}
