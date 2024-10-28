package state

import (
	"errors"
	"log"

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
	return id.String()[:5]
}

func PlayerReady(playerId string, ready bool) error {
	player := players[playerId]
	if player == nil {
		return errors.New("player not found")
	}

	game := games[player.GameId]
	if game == nil {
		return errors.New("game not found")
	}

	if game.Opponent == nil {
		return errors.New("waiting for opponent")
	}

	player.Ready = ready
	log.Println("Player", playerId, "is ready:", ready)
	return nil
}

func StartGame(hostId string) (*Game, error) {
	player := players[hostId]
	if player == nil {
		return nil, errors.New("player not found")
	}

	game := games[player.GameId]
	if game == nil {
		return nil, errors.New("game not found")
	}
	if game.State != Ready {
		return nil, errors.New("waiting for opponent")
	}
	if !game.Owner.Ready || !game.Opponent.Ready {
		return nil, errors.New("players are not ready")
	}
	if game.Owner.Id != hostId {
		return nil, errors.New("only the host can start the game")
	}

	game.State = Running
	return game, nil
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
	player := players[ownerId]
	if player == nil {
		return "", errors.New("player not found")
	}

	game := &Game{
		Id:    generateGameId(),
		Owner: players[ownerId],
		State: WaitingForOpponent,
	}
	games[game.Id] = game
	player.GameId = game.Id
	return game.Id, nil
}

func RemovePlayersGames(game *Game) {
	if game.Owner != nil {
		game.Owner.GameId = ""
	}
	if game.Opponent != nil {
		game.Opponent.GameId = ""
	}
}

func FindGame(playerId string) (*Game, error) {
	player := players[playerId]
	if player == nil {
		return nil, errors.New("player not found")
	}

	game := games[player.GameId]
	return game, nil
}

func RemoveGame(gameId string) error {
	game := games[gameId]
	if game == nil {
		return errors.New("game not found")
	}
	RemovePlayersGames(game)
	delete(games, gameId)
	return nil
}

func CloseGame(gameId string) (string, error) {
	game := games[gameId]
	RemovePlayersGames(game)
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
	game.State = Ready
	player.GameId = gameId
	return nil
}
