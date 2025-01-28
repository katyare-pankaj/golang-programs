package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func eventHandler(wg *sync.WaitGroup, eventCh <-chan int, id int) {
	defer wg.Done()
	for event := range eventCh {
		fmt.Printf("Event Handler %d: Processing event %d\n", id, event)
		// Simulate some work by sleeping for a random duration
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}
func main() {
	// Number of event handlers
	numHandlers := 3
	// Channel for sending events to handlers
	eventCh := make(chan int)
	var wg sync.WaitGroup
	// Start event handlers
	for i := 0; i < numHandlers; i++ {
		wg.Add(1)
		go eventHandler(&wg, eventCh, i+1)
	}
	// Generate some events and send them to the event handlers
	numEvents := 10
	for i := 0; i < numEvents; i++ {
		eventCh <- i + 1
	}
	// Close the event channel to signal no more events
	close(eventCh)
	// Wait for all event handlers to complete
	wg.Wait()
	fmt.Println("All event handlers have completed.")
}
