package app

import (
	"fmt"
	"time"

	"github.com/bartekkur1/cli-typeracer/contract/communication"
)

type GameLobbyScreen struct {
	ready         bool
	opponentReady bool
	hostingGame   bool
	gameStarting  bool
	startDate     time.Time
}

func (j *GameLobbyScreen) Render() {
	fmt.Printf("Game Lobby Screen\n")

	if j.gameStarting {
		if !time.Now().After(j.startDate) {
			fmt.Printf("Game starts in %s\n", time.Until(j.startDate).String())
		}
	} else {
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

func (j *GameLobbyScreen) Init(game *Game) {
	j.ready = false
	j.opponentReady = false
	j.gameStarting = false
	j.hostingGame = game.store.hostingGame
}

func (j *GameLobbyScreen) HandleEsc(game *Game) {
	game.SendMessage(communication.PlayerLeave, "")
	game.ChangeScreen(MainMenu)
}

func (j *GameLobbyScreen) GetInputHandlers(game *Game) []InputHandler {
	return []InputHandler{
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
}

func (j *GameLobbyScreen) GetNetworkHandlers(game *Game) []NetworkHandler {
	return []NetworkHandler{
		{
			event: communication.PlayerLeft,
			callback: func(e Event[communication.Message]) {
				game.ChangeScreen(MainMenu)
			},
		},
		{
			event: communication.GameClosed,
			callback: func(e Event[communication.Message]) {
				game.ChangeScreen(MainMenu)
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
				game.ChangeScreen(Race)
			},
		},
	}
}
