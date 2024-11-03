package app

import (
	"github.com/bartekkur1/cli-typeracer/contract/communication"
	"github.com/gorilla/websocket"
)

type NetworkMessage = communication.Message

type NetworkManager struct {
	eventManager EventManager[NetworkMessage]
	conn         *websocket.Conn
}

func CreateNetworkManager() *NetworkManager {
	return &NetworkManager{
		eventManager: *NewEventManager[NetworkMessage](),
	}
}

func (nm *NetworkManager) SetConnection(conn *websocket.Conn) {
	nm.conn = conn
}

func (nm *NetworkManager) SendMessage(message communication.Message) {
	nm.conn.WriteJSON(message)
}

func (nm *NetworkManager) AddListener(event communication.Command, callback Callback[communication.Message]) {
	nm.eventManager.AddListener(string(event), callback)
}

func (nm *NetworkManager) RemoveListener(event communication.Command) {
	nm.eventManager.RemoveListener(string(event))
}
