package main

import (
	"github.com/bartekkur1/cli-typeracer/client/engine"
)

func main() {
	engine := engine.CreateEngine()
	go engine.RunInputManager()
}
