package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Event represents a domain event.
type Event struct {
	ID        string
	Data      interface{}
	Timestamp time.Time
}

// EventHandler processes an event.
type EventHandler func(*Event)

// EventProcessor processes events in parallel.
type EventProcessor struct {
	handlers []EventHandler
	wg       sync.WaitGroup
}

// NewEventProcessor creates a new EventProcessor with the specified handlers.
func NewEventProcessor(handlers ...EventHandler) *EventProcessor {
	return &EventProcessor{handlers: handlers}
}

// ProcessEvents processes the given events in parallel.
func (p *EventProcessor) ProcessEvents(events []*Event) {
	// Add the number of events to the wait group to wait for all goroutines to complete.
	p.wg.Add(len(events))

	// Start a goroutine for each event to process it concurrently.
	for _, event := range events {
		go func(e *Event) {
			defer p.wg.Done()
			p.processEvent(e)
		}(event)
	}

	// Wait for all goroutines to complete processing events.
	p.wg.Wait()
}

func (p *EventProcessor) processEvent(event *Event) {
	// Simulate random processing time for each event.
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	// Apply all handlers to the event.
	for _, handler := range p.handlers {
		handler(event)
	}
}

// EventLogger logs processed events.
type EventLogger struct{}

// HandleEvent implements EventHandler to log the processed event.
func (l *EventLogger) HandleEvent(event *Event) {
	fmt.Printf("Processed event: %s\n", event.ID)
}

func main() {
	// Initialize event processor with a single event logger handler.
	processor := NewEventProcessor(&EventLogger{}.HandleEvent)

	// Generate some sample events.
	events := []*Event{
		{ID: "1", Data: "event-data-1", Timestamp: time.Now()},
		{ID: "2", Data: "event-data-2", Timestamp: time.Now()},
		{ID: "3", Data: "event-data-3", Timestamp: time.Now()},
		// Add more events as needed.
	}

	// Start processing events concurrently.
	processor.ProcessEvents(events)

	// All events will be processed before this line is reached.
	fmt.Println("All events processed.")
}
