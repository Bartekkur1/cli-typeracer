package app

var gameScreens = map[GameState]Screen{
	MainMenu: &MainMenuScreen{},
	JoinGame: &JoinGameScreen{},
	Register: &RegisterScreen{},
}
