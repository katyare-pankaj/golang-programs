package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numEvents   = 10
	numServices = 3
)

// Event represents an event in the system
type Event struct {
	ID   int
	Data string
}

// EventProcessor simulates processing an event
func EventProcessor(wg *sync.WaitGroup, events <-chan Event, serviceName string, logs chan<- string) {
	defer wg.Done()
	for event := range events {
		// Simulate some work being done by the service
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

		logs <- fmt.Sprintf("Service %s processed event %d: %s", serviceName, event.ID, event.Data)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Create a channel to pass events between the services
	events := make(chan Event)

	// Create a channel to log processed events
	logs := make(chan string)

	// WaitGroup to synchronize the event processing
	var wg sync.WaitGroup

	// Start multiple event processing services
	wg.Add(numServices)
	for i := 1; i <= numServices; i++ {
		serviceName := fmt.Sprintf("Service %d", i)
		go EventProcessor(&wg, events, serviceName, logs)
	}

	// Generate and dispatch events
	for i := 1; i <= numEvents; i++ {
		event := Event{ID: i, Data: fmt.Sprintf("Event Data %d", i)}
		events <- event
	}

	// Close the events channel to signal the end of events
	close(events)

	// Wait for all services to complete processing
	wg.Wait()

	// Consume log messages and print them
	close(logs)
	for log := range logs {
		fmt.Println(log)
	}

	fmt.Println("All events processed and logged.")
}
