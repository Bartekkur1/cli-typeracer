package main

import (
	"cli-typeracer/client/engine"
	"cli-typeracer/client/types"

	"github.com/eiannone/keyboard"
)

func main() {

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()
	var gameState = types.Engine{GameState: types.MainMenu}

	for {
		engine.PrintEngine(&gameState)
		if gameState.GameState == types.Exit {
			break
		}
	}
}
