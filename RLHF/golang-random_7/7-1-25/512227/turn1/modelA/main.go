package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numEvents  = 100000 // Number of events to generate
	numWorkers = 10     // Number of goroutines to process events
	eventRate  = 100    // Events per second
)

var wg sync.WaitGroup

// Event struct represents a simple event
type Event struct {
	ID int
}

// ProcessEvent simulates processing an event
func ProcessEvent(event Event) {
	// Simulate processing time
	time.Sleep(time.Millisecond * time.Duration(10))
	// Log event processed
	fmt.Printf("Processed event %d\n", event.ID)
	wg.Done()
}

// GenerateEvents generates events at the specified rate
func GenerateEvents() {
	for i := 1; i <= numEvents; i++ {
		event := Event{ID: i}
		// Schedule event for processing
		wg.Add(1)
		go ProcessEvent(event)

		// Wait for the specified interval before generating the next event
		time.Sleep(time.Duration(1000/eventRate) * time.Millisecond)
	}
}

func main() {
	startTime := time.Now()

	// Start generating events
	go GenerateEvents()

	// Wait for all events to be processed
	wg.Wait()

	// Calculate elapsed time and throughput
	elapsedTime := time.Since(startTime)
	throughput := float64(numEvents) / elapsedTime.Seconds()

	// Output results
	fmt.Printf("\nProcessed %d events in %.2f seconds\n", numEvents, elapsedTime.Seconds())
	fmt.Printf("Throughput: %.2f events per second\n", throughput)
}
