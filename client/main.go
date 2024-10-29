package main

import (
	"github.com/bartekkur1/cli-typeracer/client/engine"
	"github.com/bartekkur1/cli-typeracer/client/types"

	"github.com/eiannone/keyboard"
)

func main() {

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()
	var gameState = types.CreateEngine()

	for {
		engine.PrintEngine(&gameState)
		if gameState.GameState == types.Exit {
			break
		}
	}
}
