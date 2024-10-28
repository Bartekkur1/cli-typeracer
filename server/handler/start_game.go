package handler

import (
	"cli-typeracer/server/communication"
	"cli-typeracer/server/state"
	"cli-typeracer/server/util"
	"strconv"
	"time"
)

func HandleStartGame(message *communication.Message) (communication.Message, error) {

	game, err := state.StartGame(message.PlayerId)
	if err != nil {
		return communication.NewMessage(communication.Error, message.PlayerId, err.Error()), err
	}

	startDate := strconv.FormatInt(time.Now().Add(5*time.Second).UnixMilli(), 10)
	util.SendPlayerMessage(game.Owner, communication.NewMessage(communication.GameStarted, game.Owner.Id, startDate))
	util.SendPlayerMessage(game.Opponent, communication.NewMessage(communication.GameStarted, game.Owner.Id, startDate))

	return communication.NewMessage(communication.ACK, message.PlayerId, ""), nil
}
