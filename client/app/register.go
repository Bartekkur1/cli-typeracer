package app

import (
	"fmt"

	"github.com/bartekkur1/cli-typeracer/contract/communication"
	"github.com/eiannone/keyboard"
)

type RegisterScreen struct {
	registered bool
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

func (s *RegisterScreen) HandleEsc(game *Game) {
	game.Exit()
}

func (s *RegisterScreen) GetInputHandlers(game *Game) []InputHandler {
	return []InputHandler{
		{
			event: ToKey(keyboard.KeySpace),
			callback: func(e Event[KeyboardInput]) {
				if s.registered {
					game.ChangeScreen(MainMenu)
				}
			},
		},
	}
}

func (s *RegisterScreen) GetNetworkHandlers(game *Game) []NetworkHandler {
	return []NetworkHandler{
		{
			event: communication.Welcome,
			callback: func(e Event[communication.Message]) {
				s.registered = true
				game.store.playerId = e.Data.PlayerId
			},
		},
	}
}
