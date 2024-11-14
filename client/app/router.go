package app

var gameScreens = map[GameState]Screen{
	Register:  &RegisterScreen{},
	MainMenu:  &MainMenuScreen{},
	JoinGame:  &JoinGameScreen{},
	HostGame:  &HostGameScreen{},
	GameLobby: &GameLobbyScreen{},
	Race:      &RaceScreen{},
	Error:     &ErrorScreen{},
}
