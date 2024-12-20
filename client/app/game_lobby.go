package app

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bartekkur1/cli-typeracer/contract/communication"
)

// Screen Flow:
// 1. Waiting for opponent
// 2. Ready check
// 3. Game starting countdown

type GameLobbyScreen struct {
	ready           bool      // Game owner ready?
	opponentReady   bool      // Have opponent sent ready check
	hostingGame     bool      // Is player hosting the game
	gameStarting    bool      // Is game starting
	startDate       time.Time // Start date of the game
	inputHandlers   []InputHandler
	networkHandlers []NetworkHandler
}

// Renders the game starting countdown
func renderGameStarting(j *GameLobbyScreen) {
	if j.gameStarting {
		if !time.Now().After(j.startDate) {
			fmt.Printf("Game starts in %s\n", time.Until(j.startDate).Truncate(time.Second).String())
		}
	}
}

// Renders the ready check
func renderReadyCheck(j *GameLobbyScreen) {
	if !j.gameStarting {
		fmt.Printf("Press 'r' to ready up\n")

		if j.ready {
			fmt.Printf("You are \t ready\n")
		} else {
			fmt.Printf("You are \t not ready\n")
		}

		if j.opponentReady {
			fmt.Printf("Opponent \t is ready\n")
		} else {
			fmt.Printf("Opponent \t is not ready\n")
		}

		if j.ready && j.opponentReady && j.hostingGame {
			fmt.Printf("Press 's' to start the game\n")
		}
	}
}

func (j *GameLobbyScreen) Render() {
	fmt.Printf("Game Lobby Screen\n")

	renderGameStarting(j)
	renderReadyCheck(j)
}

func (j *GameLobbyScreen) Init(game *Game) {
	j.ready = false
	j.opponentReady = false
	j.gameStarting = false
	j.hostingGame = game.store.hostingGame
}

func (j *GameLobbyScreen) InitOnce(game *Game) {
	j.inputHandlers = []InputHandler{
		{
			event: "r",
			callback: func(e Event[KeyboardInput]) {
				if !j.gameStarting {
					j.ready = !j.ready
					if j.ready {
						game.SendMessage(communication.PlayerReady, "")
					} else {
						game.SendMessage(communication.PlayerNotReady, "")
					}
				}
			},
		},
		{
			event: "s",
			callback: func(e Event[KeyboardInput]) {
				if j.ready && j.opponentReady && j.hostingGame {
					game.SendMessage(communication.StartGame, "")
				}
			},
		},
	}
	j.networkHandlers = []NetworkHandler{
		{
			event: communication.PlayerLeft,
			callback: func(e Event[communication.Message]) {
				game.store.errorMessage = "Opponent left the game!"
				game.PopScreen()
				game.PushScreen(Error)
			},
		},
		{
			event: communication.GameClosed,
			callback: func(e Event[communication.Message]) {
				game.store.errorMessage = "Game was closed!"
				game.PopScreen()
				game.PushScreen(Error)
			},
		},
		{
			event: communication.PlayerReady,
			callback: func(e Event[communication.Message]) {
				if e.Data.Content != game.store.playerId {
					j.opponentReady = true
				}
			},
		},
		{
			event: communication.PlayerNotReady,
			callback: func(e Event[communication.Message]) {
				if e.Data.Content != game.store.playerId {
					j.opponentReady = false
				}
			},
		},
		{
			event: communication.GameStarting,
			callback: func(e Event[communication.Message]) {
				j.gameStarting = true
				j.startDate = time.Now().Add(5 * time.Second)
			},
		},
		{
			event: communication.GameStarted,
			callback: func(e Event[communication.Message]) {
				textNumber, error := strconv.Atoi(e.Data.Content)
				if error != nil {
					game.store.errorMessage = "Failed to parse text number"
					game.PopScreen()
					game.PushScreen(Error)
					return
				} else {
					game.store.textNumber = textNumber
				}
				game.PushScreen(Race)
			},
		},
	}
}

func (j *GameLobbyScreen) HandleEsc(game *Game) {
	game.SendMessage(communication.PlayerLeave, "")
	// Pop the game host screen and the game lobby screen
	game.PopScreen()
	game.PopScreen()
}

func (j *GameLobbyScreen) GetInputHandlers() []InputHandler {
	return j.inputHandlers
}

func (j *GameLobbyScreen) GetNetworkHandlers() []NetworkHandler {
	return j.networkHandlers
}
