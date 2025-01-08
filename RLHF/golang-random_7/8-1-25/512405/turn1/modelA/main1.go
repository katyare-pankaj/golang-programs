package main

import (
	"fmt"
	"sync"
	"time"
)

// Event represents a single event in our event-sourcing system.
type Event struct {
	Id   int
	Type string
	Data string
}

// Service processes events
type Service struct {
	name string
	mu   sync.Mutex // To prevent race conditions
}

// ProcessEvent processes a single event
func (s *Service) ProcessEvent(event *Event) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Simulate processing time
	time.Sleep(time.Duration(event.Id) * time.Millisecond)
	s.logEvent(event)
}

func (s *Service) logEvent(event *Event) {
	fmt.Printf("Service %s: Processed event with Id %d, Type: %s, Data: %s\n", s.name, event.Id, event.Type, event.Data)
}

func main() {
	var wg sync.WaitGroup

	// Create some events
	events := []*Event{
		{Id: 1, Type: "start", Data: "init"},
		{Id: 2, Type: "data", Data: "d1"},
		{Id: 3, Type: "command", Data: "c1"},
		{Id: 4, Type: "data", Data: "d2"},
		{Id: 5, Type: "command", Data: "c2"},
		{Id: 6, Type: "end", Data: "complete"},
	}

	// Create multiple services
	services := []*Service{
		{name: "Service A"},
		{name: "Service B"},
		{name: "Service C"},
	}

	// Distribute events to services
	for i, event := range events {
		service := services[i%len(services)]
		fmt.Printf("Dispatching event %d to Service %s\n", event.Id, service.name)

		wg.Add(1) // Increment WaitGroup for each event
		go func(s *Service, e *Event) {
			defer wg.Done() // Decrement WaitGroup when event is completed
			s.ProcessEvent(e)
		}(service, event) // Pass the current service and event
	}

	// Wait for all events to be processed by the services
	wg.Wait()

	fmt.Println("All events have been processed.")
}
