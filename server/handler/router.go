package handler

import (
	"cli-typeracer/server/communication"

	"github.com/gorilla/websocket"
)

// Handler function type definition
type Handler func(ws *websocket.Conn, message *communication.Message) (communication.Message, error)

// Command to handler mapping
var CommandHandlers = map[communication.Command]Handler{
	communication.Welcome:    HandleWelcome,
	communication.CreateGame: HandleCreateGame,
	communication.JoinGame:   HandleJoinGame,
}
