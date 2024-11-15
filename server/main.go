package main

import (
	"log"
	"net/http"

	"github.com/bartekkur1/cli-typeracer/server/socket"
)

func main() {
	// Define WebSocket endpoint
	http.HandleFunc("/", socket.HandleConnections)

	// @TODO: Read port from env
	// Start the server on port 8080
	log.Println("WebSocket server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
