package app

import (
	"fmt"

	"github.com/bartekkur1/cli-typeracer/contract/communication"
)

type HostGameScreen struct {
	inviteCode   *string
	playerJoined bool
}

func (j *HostGameScreen) Render() {
	fmt.Println("Game Lobby")
	if j.inviteCode != nil && !j.playerJoined {
		fmt.Printf("Invite Code: %s\n", *j.inviteCode)
	} else {
		fmt.Println("Creating game lobby...")
	}

	if j.playerJoined {
		fmt.Println("Player joined!")
	}
}

func (j *HostGameScreen) Init(game *Game) {
	game.SendMessage(communication.CreateGame, "")
}

func (j *HostGameScreen) HandleEsc(game *Game) {
	game.SendMessage(communication.PlayerLeave, "")
	game.ChangeScreen(MainMenu)
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
				j.playerJoined = true
			},
		},
	}
}
