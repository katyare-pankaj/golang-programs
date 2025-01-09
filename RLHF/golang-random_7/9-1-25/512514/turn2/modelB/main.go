package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numEvents        = 10000 // Number of events to generate
	eventChannelSize = 1000  // Buffered channel size
	numWorkers       = 10    // Number of worker goroutines
)

type Event struct {
	ID   int
	Data string
}

func main() {
	events := make(chan Event, eventChannelSize) // Buffered channel for events
	done := make(chan struct{})

	// Start the event loop
	go eventLoop(events, done)

	// Generate events and send them to the channel
	for i := 0; i < numEvents; i++ {
		events <- Event{ID: i, Data: fmt.Sprintf("Event %d", i)}
	}

	// Close the channel to signal the event loop to exit
	close(events)

	// Wait for the event loop to complete
	<-done

	fmt.Println("Event loop completed.")
}

func eventLoop(events chan Event, done chan struct{}) {
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Create a pool of worker goroutines
	for i := 0; i < numWorkers; i++ {
		go worker(i, events, &wg)
	}

	// Wait for all worker goroutines to finish
	wg.Wait()
	close(done)
}

func worker(workerID int, events chan Event, wg *sync.WaitGroup) {
	defer wg.Done()

	for event := range events {
		processEvent(workerID, event)
	}
}

func processEvent(workerID int, event Event) {
	fmt.Printf("Worker %d: Processing event ID %d: %s\n", workerID, event.ID, event.Data)

	// Simulate random processing time
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
}
