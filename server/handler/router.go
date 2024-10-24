package handler

import (
	"cli-typeracer/server/communication"
)

// Handler function type definition
type Handler func(message *communication.Message) (communication.Message, error)

// Command to handler mapping
var CommandHandlers = map[communication.Command]Handler{
	// communication.Welcome: HandleWelcome, - this is a exception because it requires a websocket to be passed as a parameter
	communication.CreateGame: HandleCreateGame,
	communication.JoinGame:   HandleJoinGame,
}
