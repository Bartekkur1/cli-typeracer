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
		fmt.Printf("User ID aquiered!\n")
		fmt.Printf("Press SPACE to continue...")
	}
}

func (s *RegisterScreen) Mount(game *Game) {
	game.networkManager.AddListener(communication.Welcome, func(message Event[communication.Message]) {
		s.registered = true
		game.store.playerId = message.Data.PlayerId
	})
	game.networkManager.SendMessage(communication.NewMessage(communication.Welcome, "", ""))

	game.inputManager.AddKeyListener(keyboard.KeySpace, func(e Event[KeyboardInput]) {
		if s.registered {
			game.ChangeScreen(MainMenu)
		}
	})
}

func (s *RegisterScreen) DisMount(game *Game) {
	game.networkManager.RemoveListener(communication.Welcome)
}

func (s *RegisterScreen) HandleEsc(game *Game) {
	game.Exit()
}
