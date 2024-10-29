package util

import (
	"os"

	"github.com/eiannone/keyboard"
)

func ReadKey() keyboard.Key {
	_, key, err := keyboard.GetKey()
	if err != nil {
		panic(err)
	}
	return key
}

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
