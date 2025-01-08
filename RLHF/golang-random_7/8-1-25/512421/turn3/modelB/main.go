package main

import (
	"fmt"
	"sync"
	"time"
)

type Event struct {
	Data string
}

func eventHandler(wg *sync.WaitGroup, eventChannel <-chan []Event) {
	defer wg.Done()
	for events := range eventChannel {
		for _, event := range events {
			fmt.Printf("Event received: %s\n", event.Data)
			time.Sleep(1 * time.Second) // Simulate some work
		}
	}
}

func main() {
	eventChannel := make(chan []Event, 100) // Buffered channel
	var wg sync.WaitGroup
	numHandlers := 3

	wg.Add(numHandlers)
	for i := 0; i < numHandlers; i++ {
		go eventHandler(&wg, eventChannel)
	}

	events := []Event{
		{Data: "Event 1"},
		{Data: "Event 2"},
		{Data: "Event 3"},
		{Data: "Event 4"},
	}

	// Batch processing
	batchSize := 2
	for i := 0; i < len(events); i += batchSize {
		end := i + batchSize
		if end > len(events) {
			end = len(events)
		}
		eventChannel <- events[i:end]
	}

	close(eventChannel)
	wg.Wait()
	fmt.Println("All events processed.")
}
