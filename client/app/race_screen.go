package app

import (
	"fmt"
	"strconv"

	"github.com/bartekkur1/cli-typeracer/client/util"
	"github.com/bartekkur1/cli-typeracer/contract/communication"
	"github.com/eiannone/keyboard"
)

// @TODO: Add race stats like WPM, accuracy, time, etc.

type RaceScreen struct {
	text             string
	progress         int
	opponentProgress int
	cursor           int
	win              bool
	finished         bool
}

// Renders give progress in the console
func renderProgress(playerName string, progress int) {
	fmt.Printf("%s [", playerName)
	for i := 0; i < 50; i++ {
		if i < progress/2 {
			fmt.Printf("▓")
		} else {
			fmt.Printf("-")
		}
	}
	fmt.Printf("] %d%%\n", progress)
}

func renderRaceText(r *RaceScreen) {
	for i, c := range r.text {
		if i < r.cursor {
			fmt.Printf("\033[32m%c\033[0m", c)
		} else if i == r.cursor {
			if c == ' ' {
				fmt.Printf("\033[31m▓\033[0m")
			} else {
				fmt.Printf("\033[31m%c\033[0m", c)
			}
		} else {
			fmt.Printf("%c", c)
		}
	}
	fmt.Println()
	fmt.Println()
}

func (r *RaceScreen) Render() {
	if r.finished {
		if r.win {
			fmt.Println("You won!")
		} else {
			fmt.Println("You lost!")
		}
		fmt.Println("Press space to continue...")
	} else {
		renderRaceText(r)
		renderProgress("You\t", r.progress)
		renderProgress("Opponent", r.opponentProgress)
	}
}

func (r *RaceScreen) Init(game *Game) {
	// @TODO: get random file from server
	r.text = util.ReadFile(1)
	r.cursor = 0
	r.progress = 0
	r.finished = false
	r.opponentProgress = 0
}

func (r *RaceScreen) HandleEsc(game *Game) {
	game.SendMessage(communication.PlayerLeave, "")
	game.PopScreen()
}

func (r *RaceScreen) GetInputHandlers(game *Game) []InputHandler {
	return []InputHandler{
		{
			event: CONSUME_ALL,
			callback: func(e Event[KeyboardInput]) {
				if r.cursor < len(r.text) && rune(e.Data.char) == rune(r.text[r.cursor]) {
					r.cursor++
					r.progress = r.cursor * 100 / len(r.text)
					if r.progress%2 == 0 {
						game.SendMessage(communication.InputProgress, fmt.Sprintf("%d", r.progress))
					}
				}
			},
		},
		{
			event: ToKey(keyboard.KeySpace),
			callback: func(e Event[KeyboardInput]) {
				if r.cursor < len(r.text) && r.text[r.cursor] == ' ' {
					r.cursor++
				} else if r.finished {
					// @TODO: Fix this, should be lobby but it's not working
					game.PopScreen()
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
				game.store.errorMessage = "Player left the game!"
				game.PopScreen()
			},
		},
		{
			event: communication.GameClosed,
			callback: func(e Event[communication.Message]) {
				game.store.errorMessage = "Game closed!"
				game.PopScreen()
			},
		},
		{
			event: communication.InputProgress,
			callback: func(e Event[communication.Message]) {
				val, err := strconv.ParseInt(e.Data.Content, 10, 32)
				if err == nil {
					r.opponentProgress = int(val)
				}
			},
		},
		{
			event: communication.GameFinished,
			callback: func(e Event[communication.Message]) {
				r.win = e.Data.Content == game.store.playerId
				r.finished = true
			},
		},
	}
}
