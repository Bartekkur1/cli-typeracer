package socket

import (
	"cli-typeracer/server/communication"
	"errors"
)

func ValidateMessage(message *communication.Message) error {
	if message.PlayerId == "" {
		return errors.New("playerId is required")
	}
	if message.Command == "" {
		return errors.New("command is required")
	}
	return nil
}
