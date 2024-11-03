package app

import (
	"fmt"

	"github.com/bartekkur1/cli-typeracer/contract/communication"
	"github.com/eiannone/keyboard"
)

type JoinGameScreen struct {
	inviteCode       string
	joinErrorMessage string
}

func (j *JoinGameScreen) Render() {
	fmt.Println("Join Game")
	if j.joinErrorMessage != "" {
		fmt.Printf("Error: %s\n", j.joinErrorMessage)
	}
	fmt.Printf("Enter invite code: %s", j.inviteCode)
}

func (j *JoinGameScreen) Init(game *Game) {
	j.inviteCode = ""
	j.joinErrorMessage = ""
}

func (j *JoinGameScreen) HandleEsc(game *Game) {
	game.ChangeScreen(MainMenu)
}

func (j *JoinGameScreen) GetInputHandlers(game *Game) []InputHandler {
	return []InputHandler{
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
}

func (j *JoinGameScreen) GetNetworkHandlers(game *Game) []NetworkHandler {
	return []NetworkHandler{
		{
			event: communication.Error,
			callback: func(e Event[communication.Message]) {
				j.joinErrorMessage = e.Data.Content
			},
		},
		{
			event: communication.GameJoined,
			callback: func(e Event[communication.Message]) {
				game.store.inviteToken = j.inviteCode
				game.ChangeScreen(HostGame)
			},
		},
	}
}
