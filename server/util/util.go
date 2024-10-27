package util

import (
	"bytes"
	"cli-typeracer/server/communication"
	"cli-typeracer/server/state"

	"github.com/gorilla/websocket"
)

func LooksLikeJSON(data []byte) bool {
	trimmed := bytes.TrimSpace(data)
	return len(trimmed) > 0 && (trimmed[0] == '{' && trimmed[len(trimmed)-1] == '}' || trimmed[0] == '[' && trimmed[len(trimmed)-1] == ']')
}

func SendPlayerMessage(player *state.Player, message communication.Message) {
	player.Conn.WriteMessage(websocket.TextMessage, []byte(communication.MessageToBytes(&message)))
}
