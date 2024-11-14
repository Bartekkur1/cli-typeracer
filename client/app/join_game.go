package app

import (
	"fmt"

	"github.com/bartekkur1/cli-typeracer/contract/communication"
	"github.com/eiannone/keyboard"
)

type JoinGameScreen struct {
	inviteCode      string
	inputHandlers   []InputHandler
	networkHandlers []NetworkHandler
}

func (j *JoinGameScreen) Render() {
	fmt.Println("Join Game")
	fmt.Printf("Enter invite code: %s", j.inviteCode)
}

func (j *JoinGameScreen) Init(game *Game) {
	j.inviteCode = ""
	game.store.hostingGame = false
}

func (j *JoinGameScreen) InitOnce(game *Game) {
	j.inputHandlers = []InputHandler{
		{
			event: CONSUME_ALL,
			callback: func(e Event[KeyboardInput]) {
				j.inviteCode += string(e.Data.char)
			},
		},
		{
			event: ToKey(keyboard.KeyBackspace2),
			callback: func(e Event[KeyboardInput]) {
				if len(j.inviteCode) > 0 {
					j.inviteCode = j.inviteCode[:len(j.inviteCode)-1]
				}
			},
		},
		{
			event: ToKey(keyboard.KeyEnter),
			callback: func(e Event[KeyboardInput]) {
				if j.inviteCode != "" {
					game.SendMessage(communication.JoinGame, j.inviteCode)
				}
			},
		},
	}
	j.networkHandlers = []NetworkHandler{
		{
			event: communication.Error,
			callback: func(e Event[communication.Message]) {
				if e.Data.Content == "game not found" {
					game.store.errorMessage = "Game not found!"
				} else {
					game.store.errorMessage = "An error occurred!"
				}
				game.PushScreen(Error)
			},
		},
		{
			event: communication.GameJoined,
			callback: func(e Event[communication.Message]) {
				game.store.inviteToken = j.inviteCode
				game.PopScreen()
				game.PushScreen(GameLobby)
			},
		},
	}
}

func (j *JoinGameScreen) HandleEsc(game *Game) {
	game.PopScreen()
}

func (j *JoinGameScreen) GetInputHandlers() []InputHandler {
	return j.inputHandlers
}

func (j *JoinGameScreen) GetNetworkHandlers() []NetworkHandler {
	return j.networkHandlers
}
