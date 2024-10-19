package screen

import (
	"cli-typeracer/client/types"
	"cli-typeracer/client/util"
)

func createMainMenuDialog() util.PickMenu {
	return util.PickMenu{Items: []types.GameState{
		types.JoinGame,
		types.HostGame,
	}, Pick: 0}
}

func RunMainMenu(engine *types.Engine) {
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
