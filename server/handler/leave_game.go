package handler

import (
	"github.com/bartekkur1/cli-typeracer/server/communication"
	"github.com/bartekkur1/cli-typeracer/server/state"
	"github.com/bartekkur1/cli-typeracer/server/util"
)

func HandlePlayerLeaveGame(message *communication.Message) (communication.Message, error) {
	game, err := state.FindGame(message.PlayerId)
	if err != nil {
		return communication.Message{}, err
	}

	if game.Owner.Id == message.PlayerId {
		util.SendPlayerMessage(game.Owner, communication.NewMessage(communication.GameClosed, game.Owner.Id, message.PlayerId))
		util.SendPlayerMessage(game.Opponent, communication.NewMessage(communication.GameClosed, game.Opponent.Id, message.PlayerId))
		state.CloseGame(game.Id)
	} else {
		util.SendPlayerMessage(game.Owner, communication.NewMessage(communication.PlayerLeft, game.Owner.Id, message.PlayerId))
		util.SendPlayerMessage(game.Opponent, communication.NewMessage(communication.PlayerLeft, game.Opponent.Id, message.PlayerId))
		state.LeaveGame(message.PlayerId)
	}

	return communication.NewMessage(communication.ACK, message.PlayerId, ""), nil
}
