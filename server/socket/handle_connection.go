package socket

import (
	"cli-typeracer/server/util"
	"log"
	"net/http"

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
		_, message, err := ws.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				log.Println("Connection closed normally:", err)
			} else {
				log.Println("Error reading message:", err)
			}
			break
		}

		if !util.LooksLikeJSON(message) {
			log.Printf("Received non-JSON message: %s\n", message)
			continue
		}

		log.Printf("Received message: %s\n", message)
		HandleMessage(ws, message)
	}
}
