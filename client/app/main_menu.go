package app

import (
	"fmt"
)

type MainMenuScreen struct {
	playerId *string
}

func (m *MainMenuScreen) Render() {
	fmt.Printf("Player ID: %s\n", *m.playerId)
	fmt.Println("Main Menu")
	fmt.Println("1. Join Game")
	fmt.Println("2. Host Game")
}

func (m *MainMenuScreen) Init(game *Game) {
	m.playerId = &game.store.playerId
}

func (m *MainMenuScreen) HandleEsc(game *Game) {
	game.Exit()
}

func (m *MainMenuScreen) GetInputHandlers(game *Game) []InputHandler {
	return []InputHandler{
		{
			event: string('1'),
			callback: func(e Event[KeyboardInput]) {
				game.ChangeScreen(JoinGame)
			},
		},
		{
			event: string('2'),
			callback: func(e Event[KeyboardInput]) {
				game.ChangeScreen(HostGame)
			},
		},
	}
}

func (m *MainMenuScreen) GetNetworkHandlers(game *Game) []NetworkHandler {
	return []NetworkHandler{}
}
