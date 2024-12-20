package handler

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/bartekkur1/cli-typeracer/contract/communication"
	"github.com/bartekkur1/cli-typeracer/server/state"
	"github.com/bartekkur1/cli-typeracer/server/util"
)

func HandleStartGame(message *communication.Message) (communication.Message, error) {
	game, err := state.StartGame(message.PlayerId)
	if err != nil {
		return communication.NewMessage(communication.Error, message.PlayerId, err.Error()), err
	}

	startDate := strconv.FormatInt(time.Now().Add(5*time.Second).UnixMilli(), 10)
	util.SendPlayerMessage(game.Owner, communication.NewMessage(communication.GameStarting, game.Owner.Id, startDate))
	util.SendPlayerMessage(game.Opponent, communication.NewMessage(communication.GameStarting, game.Owner.Id, startDate))

	time.AfterFunc(5*time.Second, func() {
		game, _ := state.FindGame(message.PlayerId)
		if game == nil {
			fmt.Printf("Game not found for %s\n", message.PlayerId)
			return
		}
		log.Printf("Game started for %s and %s", game.Owner.Id, game.Opponent.Id)
		textNumber := strconv.Itoa(util.RandomInt())
		util.SendPlayerMessage(game.Owner, communication.NewMessage(communication.GameStarted, game.Owner.Id, textNumber))
		util.SendPlayerMessage(game.Opponent, communication.NewMessage(communication.GameStarted, game.Owner.Id, textNumber))
	})

	return communication.NewMessage(communication.ACK, message.PlayerId, ""), nil
}
