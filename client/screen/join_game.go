package screen

import (
	"cli-typeracer/client/types"
	"cli-typeracer/client/util"
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

func PrintJoinGame(engine *types.Engine) {
	util.ClearConsole()
	fmt.Println("Join Game")
	fmt.Print("Enter the game code: ")

	var input string

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		if key == keyboard.KeyEsc {
			engine.GameState = types.MainMenu
			break
		}

		if key == keyboard.KeyEnter {
			util.ClearConsole()
			fmt.Printf("Joining game %s...", input)
			engine.GameState = types.Exit
			break
		}

		if char != 0 {
			fmt.Print(string(char))
			input += string(char)
		}
	}
}
