package socket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bartekkur1/cli-typeracer/contract/communication"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("Failed to upgrade connection: %v", err)
	}
	defer ws.Close()

	ws.SetCloseHandler(func(code int, text string) error {
		log.Printf("WebSocket closed: code %d, message %s\n", code, text)
		return nil
	})

	for {
		var message communication.Message
		err := ws.ReadJSON(&message)
		fmt.Printf("Received message: %+v\n", message)
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				log.Println("Connection closed normally:", err)
			} else {
				log.Println("Error reading message:", err)
				var response = communication.NewMessage(communication.Error, message.PlayerId, err.Error())
				ws.WriteMessage(websocket.TextMessage, []byte(communication.MessageToBytes(&response)))
			}
			break
		}

		log.Printf("Received message: %+v\n", message)
		HandleMessage(ws, message)
	}
}
