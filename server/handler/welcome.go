package handler

import (
	"cli-typeracer/server/communication"
	"cli-typeracer/server/state"

	"github.com/gorilla/websocket"
)

func HandleWelcome(ws *websocket.Conn) (communication.Message, error) {
	playerId := state.CreatePlayer(ws)
	return communication.NewMessage(communication.Welcome, playerId, "Welcome to the server!"), nil
}
