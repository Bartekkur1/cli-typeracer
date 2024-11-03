package app

import (
	"fmt"
	"os"
	"time"

	"github.com/bartekkur1/cli-typeracer/client/cli"
	"github.com/bartekkur1/cli-typeracer/contract/communication"
	"github.com/eiannone/keyboard"
	"github.com/gorilla/websocket"
)

type GameState int

const (
	MainMenu GameState = iota
	Register
	JoinGame
)

type RunScreen = func(*Game)

type GameStorage struct {
	exit     bool
	playerId string
}

type Game struct {
	inputManager   *InputManager
	networkManager *NetworkManager
	state          GameState
	screen         Screen
	store          GameStorage
}

func CreateGame() *Game {
	return &Game{
		inputManager:   CreateInputManager(),
		networkManager: CreateNetworkManager(),
	}
}

func (game *Game) Exit() {
	cli.ClearConsole()
	fmt.Print("Exiting...\n")
	fmt.Print("See you again!\n")
	os.Exit(0)
}

func (game *Game) StartInputManager() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	for {
		char, key, err := game.inputManager.ReadKey()
		if err != nil {
			break
		}

		if key == 0 {
			game.inputManager.EmitInput(char)
			game.inputManager.EmitCharEvent(char)
		} else {
			game.inputManager.EmitKeyEvent(key)
		}

	}
}

func (game *Game) ListenForNetwork() {
	defer game.networkManager.conn.Close()
	for {
		var message communication.Message
		err := game.networkManager.conn.ReadJSON(&message)
		if err != nil {
			panic(err)
		}

		game.networkManager.eventManager.EmitEvent(string(message.Command), message)
	}
}

// @TODO: Handle network connection failure
func (game *Game) StartServerConnection() {
	url := "ws://localhost:8080/ws"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		panic("Error connecting to WebSocket server: " + err.Error())
	}

	game.networkManager.SetConnection(conn)
	go game.ListenForNetwork()
}

func (game *Game) ChangeScreen(state GameState) {
	newScreen := gameScreens[state]
	if game.screen != nil {
		game.screen.DisMount(game)
	}
	newScreen.Mount(game)
	game.screen = newScreen
	game.state = state
}

func (game *Game) Run() {
	go game.StartInputManager()
	game.StartServerConnection()
	game.ChangeScreen(Register)

	game.inputManager.AddKeyListener(keyboard.KeyEsc, func(e InputManagerEvent) {
		game.screen.HandleEsc(game)
	})

	for {
		if game.store.exit {
			game.Exit()
		}

		cli.ClearConsole()
		game.screen.Render()
		// 60 FPS?
		time.Sleep(time.Second / 30)
	}
}
