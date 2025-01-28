package main

import (
	"fmt"
	"sync"
	"time"
)

// Event represents a simple event with a message.
type Event struct {
	message string
}

// EventLoop simulates an event loop that distributes events to handlers.
func EventLoop(events []Event, handlerCount int) {
	// Create a WaitGroup to synchronize the completion of event handlers.
	var wg sync.WaitGroup

	// Create a channel to distribute events to handlers.
	eventChan := make(chan Event)

	// Start the event handlers.
	for i := 0; i < handlerCount; i++ {
		wg.Add(1)
		go eventHandler(i, eventChan, &wg)
	}

	// Send events to the channel.
	go func() {
		for _, event := range events {
			eventChan <- event
		}
		close(eventChan) // Close the channel to signal the handlers to stop.
	}()

	// Wait for all handlers to finish processing.
	wg.Wait()
	fmt.Println("All event handlers have completed.")
}

// eventHandler processes events received from the event channel.
func eventHandler(id int, eventChan <-chan Event, wg *sync.WaitGroup) {
	defer wg.Done() // Signal completion of this handler when the function exits.

	for event := range eventChan {
		// Simulate processing the event.
		fmt.Printf("Handler %d processing event: %s\n", id, event.message)
		time.Sleep(1 * time.Second) // Simulate some work.
	}

	fmt.Printf("Handler %d has completed processing.\n", id)
}

func main() {
	// Simulate a list of events.
	events := []Event{
		{"event1"},
		{"event2"},
		{"event3"},
		{"event4"},
		{"event5"},
	}

	// Number of concurrent handlers.
	handlerCount := 3

	// Start the event loop.
	EventLoop(events, handlerCount)
}
