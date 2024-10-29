package engine

import (
	"fmt"

	"github.com/bartekkur1/cli-typeracer/client/screen"
	"github.com/bartekkur1/cli-typeracer/client/types"
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
