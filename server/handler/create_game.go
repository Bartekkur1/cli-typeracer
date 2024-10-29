package handler

import (
	"log"

	"github.com/bartekkur1/cli-typeracer/contract/communication"
	"github.com/bartekkur1/cli-typeracer/server/state"
)

func HandleCreateGame(message *communication.Message) (communication.Message, error) {
	var gameId, err = state.CreateGame(message.PlayerId)

	if err != nil {
		return communication.Message{}, err
	}

	log.Printf("Game created: %s", gameId)
	return communication.NewMessage(communication.GameCreated, message.PlayerId, gameId), nil
}
