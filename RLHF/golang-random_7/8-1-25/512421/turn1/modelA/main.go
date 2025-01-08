package main

import (
	"fmt"
	"sync"
	"time"
)

type Event struct {
	ID   int
	Data interface{}
}

func processEvent(event Event, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Processing event: %d, Data: %v\n", event.ID, event.Data)
	time.Sleep(time.Duration(event.ID) * time.Millisecond)
}

func main() {
	var wg sync.WaitGroup
	eventChan := make(chan Event)

	// Start multiple goroutines to process events
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for event := range eventChan {
				processEvent(event, &wg)
			}
		}()
	}

	// Dispatch some events
	for i := 1; i <= 10; i++ {
		eventChan <- Event{ID: i, Data: fmt.Sprintf("Event %d data", i)}
	}

	// Close the event channel when done dispatching events
	close(eventChan)

	// Wait for all goroutines to finish processing events
	wg.Wait()

	fmt.Println("All events processed")
}
