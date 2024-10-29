package socket

import (
	"errors"

	"github.com/bartekkur1/cli-typeracer/contract/communication"
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
