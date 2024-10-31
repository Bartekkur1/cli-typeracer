package engine

import (
	"github.com/bartekkur1/cli-typeracer/client/engine/ui"
	"github.com/eiannone/keyboard"
	"github.com/gorilla/websocket"
)

type CLI interface {
	Render()
	Create()
}

var stateCLI = map[GameState]CLI{
	MainMenu: &ui.MainMenu{},
}

type GameState int // GameState type

const (
	MainMenu GameState = iota
	JoinGame
	HostGame
	SetUsername
	SetServerURL
	Loading
	Exit
)

type Engine struct {
	GameState    GameState
	GameCode     string
	Socket       *websocket.Conn
	InputManager InputManager
}

// func initWebsocket() *websocket.Conn {
// 	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}
// 	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
// 	if err != nil {
// 		log.Fatal("Error connecting to WebSocket server:", err)
// 	}

// 	return conn
// }

func CreateEngine() Engine {
	// conn := initWebsocket()
	// defer conn.Close()

	return Engine{
		GameState: MainMenu,
		// Socket:    conn,
	}
}

func (engine *Engine) Render() {
	stateCLI[engine.GameState].Render()
}

func (engine *Engine) RunInputManager() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	inputManager := CreateInputManager()

	for {
		char, key, err := inputManager.ReadKey()
		if err != nil {
			break
		}

		if key == 0 {
			inputManager.EmitCharEvent(char)
		} else {
			inputManager.EmitKeyEvent(key)
		}
		if key == keyboard.KeyEsc {
			engine.GameState = Exit
			break
		}
	}
}
