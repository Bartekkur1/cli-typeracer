package main

import (
	"cli-typeracer/server/socket"
	"log"
	"net/http"
)

func main() {
	// Define WebSocket endpoint
	http.HandleFunc("/", socket.HandleConnections)

	// Start the server on port 8080
	log.Println("WebSocket server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
