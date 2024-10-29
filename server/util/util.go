package util

import (
	"bytes"

	"github.com/bartekkur1/cli-typeracer/contract/communication"
	"github.com/bartekkur1/cli-typeracer/server/state"

	"github.com/gorilla/websocket"
)

func LooksLikeJSON(data []byte) bool {
	trimmed := bytes.TrimSpace(data)
	return len(trimmed) > 0 && (trimmed[0] == '{' && trimmed[len(trimmed)-1] == '}' || trimmed[0] == '[' && trimmed[len(trimmed)-1] == ']')
}

func SendPlayerMessage(player *state.Player, message communication.Message) {
	if player.Conn == nil {
		return
	}
	player.Conn.WriteMessage(websocket.TextMessage, []byte(communication.MessageToBytes(&message)))
}
