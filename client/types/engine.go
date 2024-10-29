package types

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

type Engine struct {
	GameState GameState
	GameCode  string
	Socket    *websocket.Conn
}

func CreateEngine() Engine {
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket server:", err)
	}
	defer conn.Close()

	return Engine{
		GameState: MainMenu,
		Socket:    conn,
	}
}
