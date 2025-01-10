package main

import (
	"fmt"
	"sync"
	"time"
)

type Event struct {
	Type string
	Data interface{}
}

// eventLoop processes events from a channel using worker goroutines.
func eventLoop(events chan Event, done chan struct{}) {
	const numWorkers = 3 // Number of worker goroutines
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Create worker goroutines
	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			defer wg.Done()
			for event := range events {
				processEvent(workerID, event)
			}
		}(i)
	}

	// Wait for all worker goroutines to finish
	wg.Wait()
	close(done)
}

// processEvent handles a single event in a worker goroutine.
func processEvent(workerID int, event Event) {
	fmt.Printf("Worker %d: Processing event of type %s\n", workerID, event.Type)
	// Simulate processing time
	time.Sleep(time.Duration(100*time.Millisecond + time.Duration(workerID)*time.Second))
}

func main() {
	events := make(chan Event, 100) // Buffered channel for events
	done := make(chan struct{})

	go eventLoop(events, done)

	// Generate some events
	for i := 0; i < 10; i++ {
		events <- Event{Type: "data", Data: i}
	}

	// Signal the event loop to exit
	close(events)

	// Wait for the event loop to complete
	<-done
	fmt.Println("Event loop completed.")
}
