package app

import (
	"fmt"
)

type MainMenuScreen struct{}

func (m *MainMenuScreen) Render() {
	fmt.Println("Main Menu")
	fmt.Println("1. Join Game")
	fmt.Println("2. Host Game")
}

func (m *MainMenuScreen) Mount(game *Game) {
	game.inputManager.AddCharListener('1', func(e Event[KeyboardInput]) {
		game.ChangeScreen(JoinGame)
	})
}

func (m *MainMenuScreen) DisMount(game *Game) {
	game.inputManager.RemoveCharListener('1')
}

func (m *MainMenuScreen) HandleEsc(game *Game) {
	game.Exit()
}
