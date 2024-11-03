package app

import (
	"fmt"
)

type JoinGameScreen struct {
	inviteCode string
}

func (j *JoinGameScreen) Render() {
	fmt.Println("Join Game")
	fmt.Printf("Enter invite code: %s", j.inviteCode)
}

func (j *JoinGameScreen) Mount(game *Game) {
	j.inviteCode = ""
	game.inputManager.ListenForAll(func(e InputManagerEvent) {
		j.inviteCode += string(e.Data.char)
	})
}

func (j *JoinGameScreen) DisMount(game *Game) {
	game.inputManager.StopListeningForAll()
}

func (j *JoinGameScreen) HandleEsc(game *Game) {
	game.ChangeScreen(MainMenu)
}
