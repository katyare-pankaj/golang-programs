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
	name   string
	events []*Event
}

// ProcessEvents processes a given slice of events
func (s *Service) ProcessEvents(events []*Event) {
	s.events = events
	s.process()
}

func (s *Service) process() {
	for _, event := range s.events {
		time.Sleep(time.Duration(event.Id) * time.Millisecond) // Simulate time taken to process an event
		s.logEvent(event)
	}
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
		if i%len(services) == 0 {
			fmt.Printf("Dispatching event %d to Service %s\n", event.Id, services[i].name)
		}
		events[i].Id = i // For demonstration purposes, using index as id
		services[i%len(services)].ProcessEvents([]*Event{event})
		wg.Add(1) // Increment WaitGroup for each event started

		go func() {
			services[i%len(services)].process()
			wg.Done() // Decrement WaitGroup when event is completed
		}()
	}

	// Wait for all events to be processed by the services
	wg.Wait()

	fmt.Println("All events have been processed.")
}
