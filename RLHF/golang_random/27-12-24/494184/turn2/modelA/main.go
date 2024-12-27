package main

import (
	"errors"
	"log"
	"sync"
	"time"
)

type Event struct {
	Type string
	Data interface{}
}

type EventHandler func(Event) error

func processEvent(event Event, handler EventHandler) error {
	log.Printf("Processing event: %v\n", event)
	return handler(event)
}

func errorEventHandler(event Event) error {
	log.Printf("Error event: %v\n", event)
	// Simulate an error handling operation
	if _, ok := event.Data.(string); !ok {
		return errors.New("Invalid error event data")
	}
	return nil
}

func successEventHandler(event Event) error {
	log.Printf("Success event: %v\n", event)
	// Simulate a successful operation
	return nil
}

func main() {
	eventQueue := make(chan Event)
	var wg sync.WaitGroup

	// Start event handlers in goroutines
	wg.Add(2)
	go func() {
		for event := range eventQueue {
			if err := processEvent(event, errorEventHandler); err != nil {
				log.Printf("Error processing event: %v\n", err)
			}
		}
		wg.Done()
	}()

	go func() {
		for event := range eventQueue {
			if err := processEvent(event, successEventHandler); err != nil {
				log.Printf("Error processing event: %v\n", err)
			}
		}
		wg.Done()
	}()

	// Simulate producing events
	time.AfterFunc(time.Second, func() {
		eventQueue <- Event{Type: "error", Data: "Invalid data"}
	})

	time.AfterFunc(2*time.Second, func() {
		eventQueue <- Event{Type: "success", Data: "Operation completed"}
	})

	time.AfterFunc(3*time.Second, func() {
		eventQueue <- Event{Type: "error", Data: 42} // This will cause an error
	})

	// Wait for all event handlers to finish
	wg.Wait()
	close(eventQueue)
}
