package app

import (
	"fmt"

	"github.com/bartekkur1/cli-typeracer/client/util"
	"github.com/bartekkur1/cli-typeracer/contract/communication"
	"github.com/eiannone/keyboard"
)

type RaceScreen struct {
	text   string
	cursor int
}

func (r *RaceScreen) Render() {
	for i, c := range r.text {
		if i < r.cursor {
			// Print in green
			fmt.Printf("\033[32m%c\033[0m", c)
		} else if i == r.cursor {
			// Print in red
			fmt.Printf("\033[31m%c\033[0m", c)
		} else {
			// Print in default
			fmt.Printf("%c", c)
		}
	}
}

func (r *RaceScreen) Init(game *Game) {
	r.text = util.ReadFile(1)
	r.cursor = 0
}

func (r *RaceScreen) HandleEsc(game *Game) {
	game.SendMessage(communication.PlayerLeave, "")
	game.ChangeScreen(MainMenu)
}

func (r *RaceScreen) GetInputHandlers(game *Game) []InputHandler {
	return []InputHandler{
		{
			event: CONSUME_ALL,
			callback: func(e Event[KeyboardInput]) {
				if rune(e.Data.char) == rune(r.text[r.cursor]) {
					r.cursor++
				}
			},
		},
		{
			event: ToKey(keyboard.KeySpace),
			callback: func(e Event[KeyboardInput]) {
				if r.text[r.cursor] == ' ' {
					r.cursor++
				}
			},
		},
	}
}

func (r *RaceScreen) GetNetworkHandlers(game *Game) []NetworkHandler {
	return []NetworkHandler{
		{
			event: communication.PlayerLeft,
			callback: func(e Event[communication.Message]) {
				game.ChangeScreen(HostGame)
			},
		},
		{
			event: communication.GameClosed,
			callback: func(e Event[communication.Message]) {
				game.ChangeScreen(MainMenu)
			},
		},
	}
}
