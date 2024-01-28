package events

type Event interface{}
type EventHandler = func(event Event)
type EventManager struct {
  counter int64
  handlers map[int64]EventHandler
}

func NewEventManager() *EventManager {
  return &EventManager{
    handlers: make(map[int64]func(event Event)),
  }
}

func (m *EventManager) AddHandler(handler EventHandler) func() {
  idx := m.nextCounter()
  m.handlers[idx] = handler

  return func() {
    delete(m.handlers, idx)
  }
}

func (m *EventManager) Emit(event Event) {
  for _, handle := range m.handlers {
    handle(event)
  }
}

func (m *EventManager) nextCounter() int64 {
  counter := m.counter
  m.counter++

  return counter
}

