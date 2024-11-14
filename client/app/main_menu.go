package app

import (
	"fmt"
)

type MainMenuScreen struct {
	playerId        *string
	inputHandlers   []InputHandler
	networkHandlers []NetworkHandler
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

func (m *MainMenuScreen) InitOnce(game *Game) {
	m.inputHandlers = []InputHandler{
		{
			event: string('1'),
			callback: func(e Event[KeyboardInput]) {
				game.PushScreen(JoinGame)
			},
		},
		{
			event: string('2'),
			callback: func(e Event[KeyboardInput]) {
				game.PushScreen(HostGame)
			},
		},
	}
}

func (m *MainMenuScreen) HandleEsc(game *Game) {
	game.PopScreen()
}

func (m *MainMenuScreen) GetInputHandlers() []InputHandler {
	return m.inputHandlers
}

func (m *MainMenuScreen) GetNetworkHandlers() []NetworkHandler {
	return m.networkHandlers
}
