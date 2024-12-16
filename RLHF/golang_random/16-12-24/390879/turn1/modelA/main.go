package main

import (
	"fmt"
	"reflect"
	"sync"
)

// Event interface for all events
type Event interface{}

// UserCreated event
type UserCreated struct {
	ID   string
	Name string
}

// UserUpdated event
type UserUpdated struct {
	ID   string
	Name string
}

// EventHandler function type
type EventHandler func(Event)

// UserAggregate that holds the current state
type UserAggregate struct {
	ID   string
	Name string
}

// Apply method that applies events to the aggregate
func (u *UserAggregate) Apply(event Event) {
	eventType := reflect.TypeOf(event).Name()

	switch eventType {
	case "UserCreated":
		u.applyUserCreated(event.(UserCreated))
	case "UserUpdated":
		u.applyUserUpdated(event.(UserUpdated))
	}
}

func (u *UserAggregate) applyUserCreated(event UserCreated) {
	u.ID = event.ID
	u.Name = event.Name
	fmt.Printf("User created: %+v\n", *u)
}

func (u *UserAggregate) applyUserUpdated(event UserUpdated) {
	u.Name = event.Name
	fmt.Printf("User updated: %+v\n", *u)
}

// EventStore to store events and process them
type EventStore struct {
	events []Event
	mu     sync.Mutex
}

// AddEvent to the event store
func (es *EventStore) AddEvent(event Event) {
	es.mu.Lock()
	defer es.mu.Unlock()
	es.events = append(es.events, event)
}

// ProcessEvents processes events concurrently
func (es *EventStore) ProcessEvents(aggregate *UserAggregate) {
	var wg sync.WaitGroup

	for _, event := range es.events {
		wg.Add(1)
		go func(e Event) {
			defer wg.Done()
			aggregate.Apply(e)
		}(event)
	}

	wg.Wait()
}

func main() {
	eventStore := &EventStore{}

	// Creating events
	event1 := UserCreated{ID: "1", Name: "Alice"}
	event2 := UserUpdated{ID: "1", Name: "Alice Smith"}

	// Adding events to the store
	eventStore.AddEvent(event1)
	eventStore.AddEvent(event2)

	// Create an aggregate for User
	userAggregate := &UserAggregate{}

	// Process events concurrently
	eventStore.ProcessEvents(userAggregate)

	// Final state of the user aggregate
	fmt.Printf("Final User State: %+v\n", *userAggregate)
}
