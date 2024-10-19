package screen

import (
	"cli-typeracer/client/types"
	"cli-typeracer/client/util"
	"fmt"
	"time"
)

func PrintHostGame(engine *types.Engine) {
	util.ClearConsole()
	fmt.Println("Connecting to server...")
	time.Sleep(2 * time.Second)
	util.ClearConsole()
	fmt.Println("Generating game code...")
	time.Sleep(2 * time.Second)
	util.ClearConsole()

	fmt.Println("Your game invite code: 123123123")
	fmt.Println("Waiting for players to join...")
	time.Sleep(2 * time.Second)

	engine.GameState = types.Exit
}
