package util

import (
	"fmt"
	"os"
)

func ReadFile(fileId int) string {
	data, err := os.ReadFile("./text/" + fmt.Sprintf("%d", fileId) + ".txt")
	if err != nil {
		panic(err)
	}
	return string(data)
}

func SetTerminalSize(width, height int) {
	fmt.Printf("\x1b[8;%d;%dt", height, width)
}
