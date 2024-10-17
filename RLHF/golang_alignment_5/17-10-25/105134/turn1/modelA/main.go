package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Define an Event type to represent external events
type Event struct {
	Data string
}

// Define a Disruptor interface to abstract the source of external events
type Disruptor interface {
	Start()
	Stop()
	RegisterHandler(chan<- Event)
}

// Simulate an external event disruptor using a goroutine
type randomDisruptor struct {
	handlers []chan<- Event
	quit     chan bool
	wg       sync.WaitGroup
}

func (d *randomDisruptor) Start() {
	d.wg.Add(1)
	go func() {
		defer d.wg.Done()
		for {
			select {
			case <-d.quit:
				return
			default:
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
				event := Event{Data: fmt.Sprintf("Event %d", rand.Intn(100))}
				for _, handler := range d.handlers {
					handler <- event
				}
			}
		}
	}()
}

func (d *randomDisruptor) Stop() {
	close(d.quit)
	d.wg.Wait()
}

func (d *randomDisruptor) RegisterHandler(handler chan<- Event) {
	d.handlers = append(d.handlers, handler)
}

// API handler that processes events asynchronously
func apiHandler(events <-chan Event, wg *sync.WaitGroup) {
	defer wg.Done()
	for event := range events {
		fmt.Printf("Received event: %s\n", event.Data)
		// Perform API-related processing here
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	// Create the disruptor
	disruptor := &randomDisruptor{
		quit: make(chan bool),
	}

	// Create a channel to receive events
	eventChannel := make(chan Event)

	// Register the event handler
	disruptor.RegisterHandler(eventChannel)

	var wg sync.WaitGroup
	// Start the API handler goroutine
	wg.Add(1)
	go apiHandler(eventChannel, &wg)

	// Start the disruptor
	disruptor.Start()

	fmt.Println("API server started...")

	// Simulate application shutdown
	time.Sleep(2 * time.Second)

	// Stop the disruptor and wait for goroutines to finish
	disruptor.Stop()
	wg.Wait()

	fmt.Println("API server shutdown complete.")
}
