package event_bus

type Handler func(arg interface{})

type EventBus struct {
	events map[string][]Handler
}

func NewEventBus() *EventBus {
	return &EventBus{events: make(map[string][]Handler)}
}

func (b EventBus) On(topic string, handler Handler) {
	if arr, ok := b.events[topic]; ok {
		b.events[topic] = append(arr, handler)
	} else {
		b.events[topic] = []Handler{handler}
	}
}

func (b EventBus) Emit(topic string, arg interface{}) {
	if arr, ok := b.events[topic]; ok && len(arr) > 0 {
		for _, handler := range arr {
			handler(arg)
		}
	}
}

func (b EventBus) Off(topic string) {
	delete(b.events, topic)
}
