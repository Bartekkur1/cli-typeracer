package handler

import (
	"cli-typeracer/server/communication"
	"cli-typeracer/server/state"
)

func HandleCreateGame(message *communication.Message) (communication.Message, error) {
	var gameId, err = state.CreateGame(message.PlayerId)

	if err != nil {
		return communication.Message{}, err
	}

	return communication.NewMessage(communication.GameCreated, message.PlayerId, gameId), nil
}
