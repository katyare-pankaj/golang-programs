package main

import (
	"fmt"
	"sync"
)

type Event struct {
	data interface{}
}

var wg sync.WaitGroup
var eventQueue []Event
var mu sync.Mutex

func processEvent(event Event) {
	defer wg.Done()

	mu.Lock()
	defer mu.Unlock()

	// Simulate processing the event
	fmt.Println("Processing event:", event.data)
}

func main() {
	// Number of worker goroutines
	numWorkers := 4

	// Add initial events
	eventQueue = append(eventQueue, Event{data: "Event 1"})
	eventQueue = append(eventQueue, Event{data: "Event 2"})

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			for {
				mu.Lock()
				if len(eventQueue) == 0 {
					mu.Unlock()
					return // No more events, exit goroutine
				}
				event := eventQueue[0]
				eventQueue = eventQueue[1:]
				mu.Unlock()

				processEvent(event)
			}
		}()
	}

	// Wait for all workers to finish
	wg.Wait()
}
