package app

import (
	"fmt"

	"github.com/bartekkur1/cli-typeracer/contract/communication"
)

type HostGameScreen struct {
	inviteCode *string
}

func (j *HostGameScreen) Render() {
	fmt.Println("Host Game")
	if j.inviteCode != nil {
		fmt.Printf("Invite Code: %s\n", *j.inviteCode)
	} else {
		fmt.Println("Creating game...")
	}

	fmt.Println("Waiting for players to join...")
}

func (j *HostGameScreen) Init(game *Game) {
	game.store.hostingGame = true
	game.SendMessage(communication.CreateGame, "")
}

func (j *HostGameScreen) HandleEsc(game *Game) {
	game.SendMessage(communication.PlayerLeave, "")
	game.PopScreen()
}

func (j *HostGameScreen) GetInputHandlers(game *Game) []InputHandler {
	return []InputHandler{}
}

func (j *HostGameScreen) GetNetworkHandlers(game *Game) []NetworkHandler {
	return []NetworkHandler{
		{
			event: communication.GameCreated,
			callback: func(e Event[communication.Message]) {
				j.inviteCode = &e.Data.Content
			},
		},
		{
			event: communication.PlayerJoined,
			callback: func(e Event[communication.Message]) {
				game.PushScreen(GameLobby)
			},
		},
	}
}
