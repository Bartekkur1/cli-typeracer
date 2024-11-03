package app

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

func ToKey(key keyboard.Key) string {
	return fmt.Sprint(key)
}

type KeyboardInput struct {
	char rune
	key  keyboard.Key
}

type InputManager struct {
	eventManager EventManager[KeyboardInput]
}

type InputManagerEvent = Event[KeyboardInput]

func CreateInputManager() *InputManager {
	return &InputManager{
		eventManager: *NewEventManager[KeyboardInput](),
	}
}

func (im *InputManager) ReadKey() (rune, keyboard.Key, error) {
	char, key, err := keyboard.GetKey()
	return char, key, err
}

func (im *InputManager) AddCharListener(char rune, callback Callback[KeyboardInput]) {
	im.eventManager.AddListener(string(char), callback)
}

func (im *InputManager) AddKeyListener(key keyboard.Key, callback Callback[KeyboardInput]) {
	im.eventManager.AddListener(ToKey(key), callback)
}

func (im *InputManager) AddListener(event string, callback Callback[KeyboardInput]) {
	im.eventManager.AddListener(event, callback)
}

func (im *InputManager) EmitCharEvent(char rune) {
	im.eventManager.EmitEvent(string(char), KeyboardInput{
		char: char,
	})
}

func (im *InputManager) EmitKeyEvent(key keyboard.Key) {
	im.eventManager.EmitEvent(ToKey(key), KeyboardInput{
		key: key,
	})
}

func (im *InputManager) RemoveCharListener(char rune) {
	im.eventManager.RemoveListener(string(char))
}

func (im *InputManager) EmitInput(char rune) {
	if im.eventManager.HasListener(CONSUME_ALL) {
		im.eventManager.EmitEvent(CONSUME_ALL, KeyboardInput{
			char: char,
		})
	}
}

func (im *InputManager) ListenForAll(callback Callback[KeyboardInput]) {
	im.eventManager.AddListener(CONSUME_ALL, callback)
}

func (im *InputManager) RemoveKeyListener(key keyboard.Key) {
	im.eventManager.RemoveListener(ToKey(key))
}
