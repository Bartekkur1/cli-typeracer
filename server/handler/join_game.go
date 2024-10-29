package handler

import (
	"github.com/bartekkur1/cli-typeracer/contract/communication"
	"github.com/bartekkur1/cli-typeracer/server/state"
	"github.com/bartekkur1/cli-typeracer/server/util"
)

func HandleJoinGame(message *communication.Message) (communication.Message, error) {
	game, err := state.JoinGame(message.Content, message.PlayerId)

	if err != nil {
		return communication.Message{}, err
	}

	if game.Owner.Id != message.PlayerId {
		util.SendPlayerMessage(game.Owner, communication.NewMessage(communication.PlayerJoined, game.Owner.Id, message.PlayerId))
	}
	return communication.NewMessage(communication.GameJoined, message.PlayerId, ""), nil
}
