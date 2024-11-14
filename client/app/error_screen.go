package app

import "fmt"

type ErrorScreen struct {
	errorMessage    *string
	inputHandlers   []InputHandler
	networkHandlers []NetworkHandler
}

func (e *ErrorScreen) Render() {
	// Print error header in red on white background
	fmt.Printf("\033[41;37mError: %s\033[0m\n", *e.errorMessage)
	fmt.Printf("An error occurred. Press ESC to continue.\n")
}

func (e *ErrorScreen) Init(game *Game) {
	e.errorMessage = &game.store.errorMessage
}

func (e *ErrorScreen) InitOnce(game *Game) {
	e.inputHandlers = []InputHandler{}
	e.networkHandlers = []NetworkHandler{}
}

func (e *ErrorScreen) HandleEsc(game *Game) {
	game.PopScreen()
}

func (e *ErrorScreen) GetInputHandlers() []InputHandler {
	return e.inputHandlers
}

func (e *ErrorScreen) GetNetworkHandlers() []NetworkHandler {
	return e.networkHandlers
}
