package app

import (
	"fmt"
	"os"
	"time"

	"github.com/bartekkur1/cli-typeracer/client/cli"
	"github.com/eiannone/keyboard"
)

type GameState int

const (
	MainMenu GameState = iota
	JoinGame
)

type RunScreen = func(*Game)

type Game struct {
	inputManager *InputManager
	state        GameState
	screen       Screen
	exit         bool
}

func CreateGame() *Game {
	return &Game{
		inputManager: CreateInputManager(),
	}
}

func (game *Game) Exit() {
	cli.ClearConsole()
	fmt.Print("Exiting...\n")
	os.Exit(0)
}

func (game *Game) StartInputManager() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	for {
		char, key, err := game.inputManager.ReadKey()
		if err != nil {
			break
		}

		if key == 0 {
			game.inputManager.EmitInput(char)
			game.inputManager.EmitCharEvent(char)
		} else {
			game.inputManager.EmitKeyEvent(key)
		}

	}
}

func (game *Game) ChangeScreen(state GameState) {
	newScreen := gameScreens[state]
	if game.screen != nil {
		game.screen.DisMount(game)
	}
	newScreen.Mount(game)
	game.screen = newScreen
	game.state = state
}

func (game *Game) Run() {
	go game.StartInputManager()
	game.ChangeScreen(MainMenu)

	game.inputManager.AddKeyListener(keyboard.KeyEsc, func(e InputManagerEvent) {
		game.screen.HandleEsc(game)
	})

	for {
		if game.exit {
			game.Exit()
		}

		cli.ClearConsole()
		game.screen.Render()
		// 60 FPS
		time.Sleep(1 * time.Second / 60)
	}
}
