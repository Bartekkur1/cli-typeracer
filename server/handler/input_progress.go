package handler

import (
	"errors"

	"github.com/bartekkur1/cli-typeracer/contract/communication"
	"github.com/bartekkur1/cli-typeracer/server/state"
	"github.com/bartekkur1/cli-typeracer/server/util"
)

func HandleInputHandler(message *communication.Message) (communication.Message, error) {
	game, err := state.FindGame(message.PlayerId)

	if err != nil {
		return communication.Message{}, err
	}

	if game.State != state.Running {
		return communication.Message{}, errors.New(("game is not running"))
	}

	if game.Owner.Id != message.PlayerId && !game.OpponentFinished {
		util.SendPlayerMessage(game.Owner,
			communication.NewMessage(
				communication.InputProgress,
				game.Opponent.Id,
				message.Content,
			))
	} else if game.Owner.Id == message.PlayerId && !game.OwnerFinished {
		util.SendPlayerMessage(game.Opponent,
			communication.NewMessage(
				communication.InputProgress,
				game.Owner.Id,
				message.Content,
			))
	}

	// @TODO: This is a temporary solution, should be moved to a server side for validation
	if message.Content == "100" {
		if game.Owner.Id == message.PlayerId {
			game.OwnerFinished = true
		} else {
			game.OpponentFinished = true
		}

		if game.Winner == nil {
			if game.OwnerFinished {
				game.Winner = game.Owner
			} else {
				game.Winner = game.Opponent
			}
		}
	}

	if game.OwnerFinished && game.OpponentFinished {
		game.OwnerFinished = false
		game.OpponentFinished = false
		game.Opponent.Ready = false
		game.Owner.Ready = false
		util.SendPlayerMessage(game.Owner, communication.NewMessage(communication.GameFinished, game.Owner.Id, game.Winner.Id))
		util.SendPlayerMessage(game.Opponent, communication.NewMessage(communication.GameFinished, game.Opponent.Id, game.Winner.Id))
		game.State = state.Running
	}

	return communication.NewMessage(communication.ACK, message.PlayerId, ""), nil
}
