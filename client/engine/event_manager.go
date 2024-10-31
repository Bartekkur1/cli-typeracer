package engine

type Event[T any] struct {
	Name string
	Data T
}

type Callback[T any] func(Event[T])

type EventManager[T any] struct {
	listeners map[string][]Callback[T]
}

func NewEventManager[T any]() *EventManager[T] {
	return &EventManager[T]{
		listeners: make(map[string][]Callback[T]),
	}
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
