package util

import (
	"fmt"
	"os"
)

var fileMap = map[int]string{
	1: "1.txt",
}

func ReadFile(fileId int) string {
	data, err := os.ReadFile("./text/" + fileMap[fileId])
	if err != nil {
		panic(err)
	}
	return string(data)
}

func SetTerminalSize(width, height int) {
	fmt.Printf("\x1b[8;%d;%dt", height, width)
}
