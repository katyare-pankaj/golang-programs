package main

import (
	"fmt"
	"sync"
)

type Event struct {
	Type string
	Data interface{}
}

func eventProcessor(wg *sync.WaitGroup, events <-chan Event) {
	defer wg.Done()

	for event := range events {
		fmt.Printf("Processing event: %v\n", event)
	}
}

func main() {
	var wg sync.WaitGroup
	events := make(chan Event)

	wg.Add(1)
	go eventProcessor(&wg, events)

	events <- Event{Type: "data_received", Data: "Hello, world!"}
	events <- Event{Type: "error_occured", Data: "Something went wrong"}

	close(events)
	wg.Wait()
}
