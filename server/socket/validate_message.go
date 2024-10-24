package socket

import (
	"cli-typeracer/server/communication"
	"errors"
)

func ValidateMessage(message *communication.Message) error {
	if message.PlayerId == "" && message.Command != communication.Welcome {
		return errors.New("playerId is required")
	}
	if message.Command == "" {
		return errors.New("command is required")
	}
	return nil
}
