package main

import (
	"github.com/bartekkur1/cli-typeracer/client/app"
)

func main() {
	game := app.CreateGame()
	game.Run()
}
