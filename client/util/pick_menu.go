package util

import (
	"cli-typeracer/client/types"
	"fmt"

	"github.com/eiannone/keyboard"
)

type PickMenu struct {
	Items []types.GameState
	Pick  int
}

func RunMenu(menu *PickMenu) int {
	for {
		ClearConsole()
		printMenu(menu)
		key := ReadKey()

		if key == keyboard.KeyArrowUp {
			menu.Pick--
			if menu.Pick < 0 {
				menu.Pick = len(menu.Items) - 1
			}
		} else if key == keyboard.KeyArrowDown {
			menu.Pick++
			if menu.Pick >= len(menu.Items) {
				menu.Pick = 0
			}
		} else if key == keyboard.KeyEnter {
			return menu.Pick + 1
		} else if key == keyboard.KeyEsc {
			break
		}
	}

	return 0
}

func ClearConsole() {
	fmt.Print("\033[H\033[2J")
}

func printMenu(menu *PickMenu) {
	fmt.Println("Welcome to Typeracer!")
	for i, item := range menu.Items {
		if menu.Pick == i {
			// Print item with red background
			fmt.Printf("\033[41m%d: %s\033[0m\n", i+1, item)
		} else {
			fmt.Printf("%d: %s\n", i+1, item)
		}
	}
}
