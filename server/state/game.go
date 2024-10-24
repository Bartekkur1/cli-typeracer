package state

import "github.com/gorilla/websocket"

type GameState string

const (
	WaitingForOpponent GameState = "WaitingForOpponent"
	Running            GameState = "Running"
	Finished           GameState = "Finished"
)

type Player struct {
	Id     string
	GameId string
	Conn   *websocket.Conn
}

type Game struct {
	Id       string
	State    GameState
	Owner    *Player
	Opponent *Player
}