package app

import (
	"fmt"

	"github.com/bartekkur1/cli-typeracer/contract/communication"
)

type HostGameScreen struct {
	inviteCode      *string
	inputHandlers   []InputHandler
	networkHandlers []NetworkHandler
}

func (h *HostGameScreen) Render() {
	fmt.Println("Host Game")
	if h.inviteCode != nil {
		fmt.Printf("Invite Code: %s\n", *h.inviteCode)
	} else {
		fmt.Println("Creating game...")
	}

	fmt.Println("Waiting for players to join...")
}

func (h *HostGameScreen) Init(game *Game) {
	game.store.hostingGame = true
	game.SendMessage(communication.CreateGame, "")
}

func (h *HostGameScreen) InitOnce(game *Game) {
	h.inputHandlers = []InputHandler{}
	h.networkHandlers = []NetworkHandler{
		{
			event: communication.GameCreated,
			callback: func(e Event[communication.Message]) {
				h.inviteCode = &e.Data.Content
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

func (h *HostGameScreen) HandleEsc(game *Game) {
	game.SendMessage(communication.PlayerLeave, "")
	game.PopScreen()
}

func (h *HostGameScreen) GetInputHandlers() []InputHandler {
	return h.inputHandlers
}

func (h *HostGameScreen) GetNetworkHandlers() []NetworkHandler {
	return h.networkHandlers
}
