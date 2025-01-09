package main

import (
	"fmt"
	"sync"
	"time"
)

// Event represents an incoming event.
type Event struct {
	ID   int
	Data interface{}
	Time time.Time
}

// processEvent simulates processing an event.
func processEvent(event Event) {
	fmt.Printf("Processing event: ID=%d, Data=%v, Time=%v\n", event.ID, event.Data, event.Time)
	// Simulate processing time
	time.Sleep(time.Duration(100+event.ID%100) * time.Millisecond)
}

// worker processes events from the event channel.
func worker(id int, eventChannel chan Event, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case event, ok := <-eventChannel:
			if !ok {
				fmt.Printf("Worker %d: Channel closed, exiting\n", id)
				return
			}
			processEvent(event)
		default:
			time.Sleep(1 * time.Millisecond)
		}
	}
}

func main() {
	const numWorkers = 10  // Number of worker goroutines
	const bufferSize = 100 // Buffer size for the event channel

	eventChannel := make(chan Event, bufferSize)
	wg := &sync.WaitGroup{}

	// Create worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, eventChannel, wg)
	}

	// Generate events and send them to the channel
	for i := 0; i < 1000; i++ {
		event := Event{ID: i, Data: fmt.Sprintf("Data %d", i), Time: time.Now()}
		eventChannel <- event
	}

	// Close the channel to signal workers to exit
	close(eventChannel)

	// Wait for all worker goroutines to finish
	wg.Wait()
	fmt.Println("All events processed.")
}
