package util

import (
	"github.com/eiannone/keyboard"
)

func ReadKey() keyboard.Key {
	_, key, err := keyboard.GetKey()
	if err != nil {
		panic(err)
	}
	return key
}
