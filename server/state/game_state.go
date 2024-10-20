package state

import (
	"github.com/gorilla/websocket"
)

type Player struct {
	Id         string
	Connection *websocket.Conn
	GameId     string
}

type GameState struct {
	Id       string
	Owner    *Player
	Opponent *Player
}

var games = make(map[string]*GameState)

func generateGameId() string {
	return "1"
}

func CreateGame(owner *Player) (string, error) {
	gameId := generateGameId()
	game := &GameState{
		Id: gameId,
	}
	owner.GameId = gameId
	games[gameId] = game
	game.Owner = owner
	return gameId, nil
}

func DisconnectPlayers(game *GameState) {
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

func JoinGame(gameId string, player *Player) (message string, err error) {
	game := games[gameId]
	if game == nil {
		return "Game not found", nil
	}
	game.Opponent = player
	player.GameId = gameId
	return "Player joined", nil
}
