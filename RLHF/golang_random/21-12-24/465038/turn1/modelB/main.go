package main

import (
	"fmt"
	"sync"
	"time"
)

// EventSourcingAggregate is an interface for aggregates that handle events.
type EventSourcingAggregate interface {
	// HandleEvent applies an event to the aggregate's state.
	HandleEvent(interface{})
}

// SimpleAggregate is a basic aggregate that counts events.
type SimpleAggregate struct {
	eventsCount int
}

// HandleEvent updates the count of events in the aggregate.
func (a *SimpleAggregate) HandleEvent(e interface{}) {
	a.eventsCount++
}

// EventPublisher is responsible for publishing events to the aggregates.
type EventPublisher struct {
	aggregates map[string]EventSourcingAggregate
}

// NewEventPublisher creates an EventPublisher instance with the provided aggregates.
func NewEventPublisher(aggregates map[string]EventSourcingAggregate) *EventPublisher {
	return &EventPublisher{
		aggregates: aggregates,
	}
}

// PublishEvent publishes an event to all registered aggregates.
func (ep *EventPublisher) PublishEvent(event interface{}) {
	// WaitGroup to synchronize the execution of aggregate handling.
	var wg sync.WaitGroup

	for _, aggregate := range ep.aggregates {
		// For each aggregate, create a new goroutine to handle the event.
		wg.Add(1)
		go func(a EventSourcingAggregate) {
			// Handle the event on the aggregate.
			a.HandleEvent(event)
			// Signal that the aggregate has finished processing the event.
			wg.Done()
		}(aggregate)
	}

	// Wait for all aggregates to finish handling the event.
	wg.Wait()
}

func main() {
	// Create a few aggregates to listen for events.
	aggregate1 := &SimpleAggregate{}
	aggregate2 := &SimpleAggregate{}
	aggregate3 := &SimpleAggregate{}

	aggregates := map[string]EventSourcingAggregate{
		"aggregate1": aggregate1,
		"aggregate2": aggregate2,
		"aggregate3": aggregate3,
	}

	eventPublisher := NewEventPublisher(aggregates)

	// Publish an event that will be handled by all aggregates.
	eventPublisher.PublishEvent("event-data")

	// Since all aggregates run in separate goroutines, we need to ensure that
	// they have enough time to process the event before reading their counts.
	time.Sleep(100 * time.Millisecond)

	// Verify that the aggregates received the event and updated their states.
	fmt.Println("Aggregate 1 event count:", aggregate1.eventsCount)
	fmt.Println("Aggregate 2 event count:", aggregate2.eventsCount)
	fmt.Println("Aggregate 3 event count:", aggregate3.eventsCount)

	// Expected output:
	// Aggregate 1 event count: 1
	// Aggregate 2 event count: 1
	// Aggregate 3 event count: 1
}
