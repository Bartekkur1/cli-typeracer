package handler

import (
	"cli-typeracer/server/communication"
	"cli-typeracer/server/state"
)

var commandToStringMap = map[communication.Command]string{
	communication.Ready:    "ready",
	communication.NotReady: "not ready",
}

func HandleReadyCheck(message *communication.Message) (communication.Message, error) {
	err := state.PlayerReady(message.PlayerId, message.Command == communication.Ready)

	if err != nil {
		return communication.Message{}, err
	}

	return communication.NewMessage(communication.ACK, message.PlayerId, "Player "+message.PlayerId+" is "+commandToStringMap[message.Command]), nil
}
