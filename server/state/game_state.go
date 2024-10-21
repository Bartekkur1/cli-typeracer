package state

import (
	"errors"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Player struct {
	Id         string
	Connection *websocket.Conn
	GameId     string
}

type GameState string

const (
	WaitingForOpponent GameState = "WaitingForOpponent"
	Running            GameState = "Running"
	Finished           GameState = "Finished"
)

type Game struct {
	Id       string
	State    GameState
	Owner    *Player
	Opponent *Player
}

var games = make(map[string]*Game)
var players = make(map[string]*Player)

func generateGameId() string {
	id := uuid.New()
	return id.String()
}

func CreatePlayer(id string, conn *websocket.Conn) {
	player := &Player{
		Id:         id,
		Connection: conn,
	}
	players[id] = player
}

func CreateGame(ownerId string) (string, error) {
	if players[ownerId] == nil {
		return "", errors.New("Player not found")
	}

	game := &Game{
		Id:    generateGameId(),
		Owner: players[ownerId],
		State: WaitingForOpponent,
	}
	games[game.Id] = game
	return game.Id, nil
}

func DisconnectPlayers(game *Game) {
	if game.Owner != nil {
		game.Owner.GameId = ""
	}
	if game.Opponent != nil {
		game.Opponent.GameId = ""
	}
}

func CloseGame(gameId string) (result string, err error) {
	game := games[gameId]
	DisconnectPlayers(game)
	if game == nil {
		return "", nil
	}
	delete(games, gameId)
	return "Game closed", nil
}

func JoinGame(gameId string, playerId string) (err error) {
	player := players[playerId]
	if player == nil {
		return errors.New("Player not found")
	}

	game := games[gameId]
	if game == nil {
		return errors.New("Game not found")
	}
	game.Opponent = player
	game.State = Running
	player.GameId = gameId
	return nil
}
