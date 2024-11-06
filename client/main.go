package main

import (
	"github.com/bartekkur1/cli-typeracer/client/app"
)

func main() {
	// util.SetTerminalSize(100, 50)
	game := app.CreateGame()
	game.Run()
}
