package main

import (
	"fmt"
	"time"
)

// Event represents a scheduled event
type Event struct {
	Name        string
	Time        time.Time
	Description string
}

// EventScheduler manages a list of events
type EventScheduler struct {
	Events []Event
}

// AddEvent adds a new event to the scheduler
func (s *EventScheduler) AddEvent(name string, t time.Time, description string) {
	s.Events = append(s.Events, Event{
		Name:        name,
		Time:        t,
		Description: description,
	})
}

// ListEvents prints all scheduled events
func (s *EventScheduler) ListEvents() {
	for _, event := range s.Events {
		fmt.Printf("Event: %s\nTime: %s\nDescription: %s\n\n", event.Name, event.Time.Format("2006-01-02 15:04:05"), event.Description)
	}
}

// Main function to demonstrate the scheduler
func main() {
	scheduler := EventScheduler{}

	// Add a couple of events
	scheduler.AddEvent("Meeting", time.Now().Add(15*time.Minute), "Weekly team meeting")
	scheduler.AddEvent("Lunch", time.Now().Add(30*time.Minute), "Lunch with colleagues")

	// List all events
	scheduler.ListEvents()
}
