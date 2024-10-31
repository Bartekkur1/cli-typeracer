package engine

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

func ToKey(key keyboard.Key) string {
	return fmt.Sprint(key)
}

type InputManager struct {
	eventManager EventManager[struct{}]
}

type InputManagerEvent = Event[struct{}]

func CreateInputManager() *InputManager {
	return &InputManager{
		eventManager: *NewEventManager[struct{}](),
	}
}

func (im *InputManager) ReadKey() (rune, keyboard.Key, error) {
	char, key, err := keyboard.GetKey()
	return char, key, err
}

func (im *InputManager) AddListener(event string, callback Callback[struct{}]) {
	im.eventManager.AddListener(event, callback)
}

func (im *InputManager) EmitCharEvent(char rune) {
	im.eventManager.EmitEvent(string(char), struct{}{})
}

func (im *InputManager) EmitKeyEvent(key keyboard.Key) {
	im.eventManager.EmitEvent(ToKey(key), struct{}{})
}
