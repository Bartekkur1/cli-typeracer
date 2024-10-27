package handler

import (
	"cli-typeracer/server/communication"
	"cli-typeracer/server/state"
	"cli-typeracer/server/util"
)

func HandleStartGame(message *communication.Message) (communication.Message, error) {

	game, err := state.StartGame(message.PlayerId)
	if err != nil {
		return communication.NewMessage(communication.Error, message.PlayerId, err.Error()), err
	}

	util.SendPlayerMessage(game.Owner, communication.NewMessage(communication.GameStarted, game.Owner.Id, "game starting in 5"))
	util.SendPlayerMessage(game.Opponent, communication.NewMessage(communication.GameStarted, game.Owner.Id, "game starting in 5"))

	return communication.NewMessage(communication.ACK, message.PlayerId, ""), nil
}
