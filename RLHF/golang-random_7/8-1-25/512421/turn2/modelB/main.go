package main

import (
	"fmt"
	"sync"
	"time"
)

type Event struct {
	Data string
}

func eventHandler(wg *sync.WaitGroup, eventChannel <-chan Event) {
	defer wg.Done()
	for event := range eventChannel {
		fmt.Printf("Event received: %s\n", event.Data)
		time.Sleep(1 * time.Second) // Simulate some work
	}
}

func main() {
	eventChannel := make(chan Event)
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

	for _, event := range events {
		eventChannel <- event
	}

	close(eventChannel)
	wg.Wait()
	fmt.Println("All events processed.")
}
