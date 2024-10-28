package handler

import (
	"cli-typeracer/server/communication"
	"cli-typeracer/server/state"
	"log"
)

func HandleCreateGame(message *communication.Message) (communication.Message, error) {
	var gameId, err = state.CreateGame(message.PlayerId)

	log.Printf("Game created: %s", gameId)

	if err != nil {
		return communication.Message{}, err
	}

	return communication.NewMessage(communication.GameCreated, message.PlayerId, gameId), nil
}
