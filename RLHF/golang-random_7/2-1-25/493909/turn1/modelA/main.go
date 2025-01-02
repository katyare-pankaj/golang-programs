package main

import (
	"fmt"
	"sync"
	"time"
)

// Event represents an event in an event-sourced domain.
type Event struct {
	ID   int
	Data string
}

// EventProcessor processes events concurrently.
type EventProcessor struct {
	events          []Event
	wg              *sync.WaitGroup
	mu              *sync.Mutex
	errors          []error
	processedEvents map[int]bool
}

// NewEventProcessor creates a new EventProcessor.
func NewEventProcessor() *EventProcessor {
	return &EventProcessor{
		wg:              &sync.WaitGroup{},
		mu:              &sync.Mutex{},
		errors:          make([]error, 0),
		processedEvents: make(map[int]bool),
	}
}

// AddEvent adds an event to the processor.
func (ep *EventProcessor) AddEvent(e Event) {
	ep.mu.Lock()
	defer ep.mu.Unlock()
	ep.events = append(ep.events, e)
}

// ProcessEvents processes all events concurrently.
func (ep *EventProcessor) ProcessEvents() {
	ep.wg.Add(len(ep.events))

	for _, event := range ep.events {
		go func(e Event) {
			defer ep.wg.Done()

			// Simulate processing time
			time.Sleep(time.Duration(100+int64(e.ID)) * time.Millisecond)

			// Simulate processing logic
			if e.ID%2 == 0 {
				// Simulate an error for even-numbered events
				ep.mu.Lock()
				ep.errors = append(ep.errors, fmt.Errorf("error processing event %d: %s", e.ID, e.Data))
				ep.mu.Unlock()
			} else {
				// Mark the event as processed
				ep.mu.Lock()
				ep.processedEvents[e.ID] = true
				ep.mu.Unlock()
				fmt.Printf("Processed event %d: %s\n", e.ID, e.Data)
			}
		}(event)
	}

	// Wait for all goroutines to finish
	ep.wg.Wait()

	// Check for any errors
	if len(ep.errors) > 0 {
		fmt.Println("Errors occurred while processing events:")
		for _, err := range ep.errors {
			fmt.Println(err)
		}
	} else {
		fmt.Println("All events processed successfully.")
	}
}

func main() {
	ep := NewEventProcessor()

	// Add some events
	for i := 1; i <= 10; i++ {
		ep.AddEvent(Event{ID: i, Data: fmt.Sprintf("Event %d", i)})
	}

	// Process events in parallel
	ep.ProcessEvents()
}
