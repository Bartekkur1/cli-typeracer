package screen

import (
	"github.com/bartekkur1/cli-typeracer/client/types"
	"github.com/bartekkur1/cli-typeracer/client/util"
)

func createMainMenuDialog() util.PickMenu {
	return util.PickMenu{Items: []types.GameState{
		types.JoinGame,
		types.HostGame,
	}, Pick: 0}
}

func PrintMainMenu(engine *types.Engine) {
	var menuDialog = createMainMenuDialog()
	var pick = util.RunMenu(&menuDialog)
	if pick == 1 {
		engine.GameState = types.JoinGame
	} else if pick == 2 {
		engine.GameState = types.HostGame
	} else if pick == 0 {
		engine.GameState = types.Exit
	}
}
