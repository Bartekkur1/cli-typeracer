package handler

import (
	"github.com/bartekkur1/cli-typeracer/contract/communication"
	"github.com/bartekkur1/cli-typeracer/server/state"

	"github.com/gorilla/websocket"
)

func HandleWelcome(ws *websocket.Conn) (communication.Message, error) {
	playerId := state.CreatePlayer(ws)
	return communication.NewMessage(communication.Welcome, playerId, "Welcome to the server!"), nil
}
