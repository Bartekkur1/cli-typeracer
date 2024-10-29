package handler

import (
	"github.com/bartekkur1/cli-typeracer/contract/communication"
	"github.com/bartekkur1/cli-typeracer/server/state"
	"github.com/bartekkur1/cli-typeracer/server/util"
)

var commandToStringMap = map[communication.Command]string{
	communication.PlayerReady:    "ready",
	communication.PlayerNotReady: "not ready",
}

func HandleReadyCheck(message *communication.Message) (communication.Message, error) {
	game, err := state.PlayerReady(message.PlayerId, message.Command == communication.PlayerReady)

	if err != nil {
		return communication.Message{}, err
	}

	util.SendPlayerMessage(game.Owner, communication.NewMessage(message.Command, game.Owner.Id, message.PlayerId))
	util.SendPlayerMessage(game.Opponent, communication.NewMessage(message.Command, game.Opponent.Id, message.PlayerId))
	return communication.NewMessage(communication.ACK, message.PlayerId, "Player "+message.PlayerId+" is "+commandToStringMap[message.Command]), nil
}
