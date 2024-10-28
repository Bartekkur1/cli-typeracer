package socket

import "github.com/gorilla/websocket"

type PlayerConnection struct {
	ID   string
	Conn *websocket.Conn
}
