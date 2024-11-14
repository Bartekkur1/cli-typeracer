package app

import (
	"fmt"

	"github.com/bartekkur1/cli-typeracer/contract/communication"
	"github.com/eiannone/keyboard"
)

type RegisterScreen struct {
	registered     bool
	inputHandler   []InputHandler
	networkHandler []NetworkHandler
}

func (s *RegisterScreen) Render() {
	fmt.Printf("Welcome to CLI TypeRacer!\n")
	fmt.Printf("Connecting to web server...\n")
	if s.registered {
		fmt.Printf("Connected to server!\n")
		fmt.Printf("User ID acquired!\n")
		fmt.Printf("Press SPACE to continue...")
	}
}

func (s *RegisterScreen) Init(game *Game) {
	game.SendMessage(communication.Welcome, "")
}

func (s *RegisterScreen) InitOnce(game *Game) {
	s.inputHandler = []InputHandler{
		{
			event: ToKey(keyboard.KeySpace),
			callback: func(e Event[KeyboardInput]) {
				if s.registered {
					game.ForceMainMenu()
				}
			},
		},
	}

	s.networkHandler = []NetworkHandler{
		{
			event: communication.Welcome,
			callback: func(e Event[communication.Message]) {
				s.registered = true
				game.store.playerId = e.Data.PlayerId
			},
		},
	}
}

func (s *RegisterScreen) HandleEsc(game *Game) {
	game.Exit()
}

func (s *RegisterScreen) GetInputHandlers() []InputHandler {
	return s.inputHandler
}

func (s *RegisterScreen) GetNetworkHandlers() []NetworkHandler {
	return s.networkHandler
}
