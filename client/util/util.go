package util

import (
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
