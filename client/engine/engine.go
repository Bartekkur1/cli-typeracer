package engine

import (
	"cli-typeracer/client/screen"
	"cli-typeracer/client/types"
	"fmt"
)

func PrintEngine(engine *types.Engine) {
	if engine.GameState == types.MainMenu {
		screen.PrintMainMenu(engine)
	} else if engine.GameState == types.JoinGame {
		screen.PrintJoinGame(engine)
	} else if engine.GameState == types.HostGame {
		screen.PrintHostGame(engine)
	} else {
		fmt.Println("Exiting...")
		engine.GameState = types.Exit
	}
}
