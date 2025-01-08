package main

import (
	"fmt"
	"sync"
	"time"
)

type Event struct {
	Type    string
	Data    interface{}
	Handler func(event Event)
}

type EventDriver struct {
	wg         sync.WaitGroup
	eventQueue chan Event
	handlers   map[string][]func(event Event)
	mu         sync.RWMutex
}

func NewEventDriver() *EventDriver {
	return &EventDriver{
		eventQueue: make(chan Event),
		handlers:   make(map[string][]func(event Event)),
	}
}

func (ed *EventDriver) Start() {
	ed.wg.Add(1)
	go ed.eventLoop()
}

func (ed *EventDriver) Stop() {
	close(ed.eventQueue)
	ed.wg.Wait()
}

func (ed *EventDriver) eventLoop() {
	defer ed.wg.Done()
	for event := range ed.eventQueue {
		ed.processEvent(event)
	}
}

func (ed *EventDriver) processEvent(event Event) {
	ed.mu.RLock()
	defer ed.mu.RUnlock()

	if handlers, exists := ed.handlers[event.Type]; exists {
		for _, handler := range handlers {
			handler(event)
		}
	}
}

func (ed *EventDriver) RegisterHandler(eventType string, handler func(event Event)) {
	ed.mu.Lock()
	defer ed.mu.Unlock()

	if _, exists := ed.handlers[eventType]; !exists {
		ed.handlers[eventType] = []func(event Event){handler}
	} else {
		ed.handlers[eventType] = append(ed.handlers[eventType], handler)
	}
}

func (ed *EventDriver) Publish(event Event) {
	ed.eventQueue <- event
}

func main() {
	ed := NewEventDriver()
	ed.Start()
	defer ed.Stop()

	// Register handlers
	ed.RegisterHandler("message", func(event Event) {
		fmt.Printf("Message handler: %v\n", event.Data)
	})

	ed.RegisterHandler("error", func(event Event) {
		fmt.Printf("Error handler: %v\n", event.Data)
	})

	// Publish events
	ed.Publish(Event{Type: "message", Data: "Hello, world!"})
	ed.Publish(Event{Type: "error", Data: "Something went wrong"})
	ed.Publish(Event{Type: "unknown", Data: "An unknown event"}) // No handler for this event type

	// For demonstration, wait for a few seconds before stopping
	fmt.Println("Waiting for event processing to complete...")
	time.Sleep(2 * time.Second)
}
