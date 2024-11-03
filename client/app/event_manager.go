package app

type Event[T any] struct {
	Name string
	Data T
}

const CONSUME_ALL = "CONSUME_ALL_CONST"

type Callback[T any] func(Event[T])

type EventManager[T any] struct {
	listeners map[string][]Callback[T]
}

func NewEventManager[T any]() *EventManager[T] {
	return &EventManager[T]{
		listeners: make(map[string][]Callback[T]),
	}
}

func (em *EventManager[T]) ListenForAll(callback Callback[T]) {
	em.listeners[CONSUME_ALL] = append(em.listeners[CONSUME_ALL], callback)
}

func (em *EventManager[T]) AddListener(event string, callback Callback[T]) {
	em.listeners[event] = append(em.listeners[event], callback)
}

func (em *EventManager[T]) EmitEvent(event string, data T) {
	if callbacks, found := em.listeners[event]; found {
		for _, callback := range callbacks {
			callback(Event[T]{Name: event, Data: data})
		}
	}
}

func (em *EventManager[T]) RemoveListener(event string) {
	if em.listeners[event] != nil {
		em.listeners[event] = []Callback[T]{}
	}
}

func (em *EventManager[T]) HasListener(event string) bool {
	return em.listeners[event] != nil
}
