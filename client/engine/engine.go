package engine

import (
	"cli-typeracer/client/screen"
	"cli-typeracer/client/types"
	"fmt"
)

func PrintEngine(engine *types.Engine) {
	if engine.GameState == types.MainMenu {
		screen.RunMainMenu(engine)
	} else if engine.GameState == types.JoinGame {
		screen.PrintJoinGame(engine)
	} else {
		fmt.Println("Exiting...")
		engine.GameState = types.Exit
	}
}
