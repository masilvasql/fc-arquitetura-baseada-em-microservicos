package events

import (
	"errors"
	"sync"
)

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

var ErrHandlerAlreadyRegistered = errors.New("Handler already registered")

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {

	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}

	ed.handlers[eventName] = append(ed.handlers[eventName], handler)

	return nil
}

func (ed *EventDispatcher) Clear() {
	ed.handlers = make(map[string][]EventHandlerInterface)
}

func (ed *EventDispatcher) Has(name string, event EventHandlerInterface) bool {
	if _, ok := ed.handlers[name]; ok {
		for _, handler := range ed.handlers[name] {
			if handler == event {
				return true
			}
		}
	}
	return false
}

func (ev *EventDispatcher) Dispatch(event EventInterface) error {
	if handlers, ok := ev.handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{}
		for _, handler := range handlers {
			wg.Add(1)
			go handler.Handle(event, wg)
		}
		wg.Wait()
	}
	return nil
}

func (ed *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	if _, ok := ed.handlers[eventName]; ok {
		for i, h := range ed.handlers[eventName] {
			if h == handler {
				ed.handlers[eventName] = append(ed.handlers[eventName][:i], ed.handlers[eventName][i+1:]...)
				return nil
			}
		}
	}
	return nil
}
