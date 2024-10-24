package handler

import (
	"cli-typeracer/server/communication"
	"cli-typeracer/server/state"
)

func HandleJoinGame(message *communication.Message) (communication.Message, error) {
	var err = state.JoinGame(message.Content, message.PlayerId)

	if err != nil {
		return communication.Message{}, err
	}

	return communication.NewMessage(communication.GameJoined, message.PlayerId, ""), nil
}
