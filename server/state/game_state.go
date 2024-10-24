package state

import (
	"errors"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var games = make(map[string]*Game)
var players = make(map[string]*Player)

func generateId() string {
	id := uuid.New()
	return id.String()
}

func generateGameId() string {
	id := uuid.New()
	return id.String()[:8]
}

func CreatePlayer(ws *websocket.Conn) string {
	id := generateId()
	player := &Player{
		Id:   id,
		Conn: ws,
	}
	players[id] = player
	return id
}

func CreateGame(ownerId string) (string, error) {
	if players[ownerId] == nil {
		return "", errors.New("player not found")
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

func CloseGame(gameId string) (string, error) {
	game := games[gameId]
	DisconnectPlayers(game)
	if game == nil {
		return "", nil
	}
	delete(games, gameId)
	return "Game closed", nil
}

func JoinGame(gameId string, playerId string) error {
	player := players[playerId]
	if player == nil {
		return errors.New("player not found")
	}

	game := games[gameId]
	if game == nil {
		return errors.New("game not found")
	}
	game.Opponent = player
	game.State = Running
	player.GameId = gameId
	return nil
}
