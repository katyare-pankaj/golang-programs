package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

// Aggregate is an abstract representation of an aggregate.
type Aggregate interface {
	ProcessEvent(event interface{})
}

// Aggregate1 is a simple example of an aggregate.
type Aggregate1 struct{}

func (a *Aggregate1) ProcessEvent(event interface{}) {
	// Simulate processing time
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	printf("Aggregate1 processed event: %v\n", event)
}

// Aggregate2 is another example of an aggregate.
type Aggregate2 struct{}

func (a *Aggregate2) ProcessEvent(event interface{}) {
	// Simulate processing time
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	printf("Aggregate2 processed event: %v\n", event)
}

func main() {
	// Create instances of the aggregates
	aggregate1 := &Aggregate1{}
	aggregate2 := &Aggregate2{}

	// Create a WaitGroup to coordinate processing
	var wg sync.WaitGroup

	// Number of events to process
	numEvents := 5

	// Process events for Aggregate1
	for i := 0; i < numEvents; i++ {
		event := fmt.Sprintf("Event %d for Aggregate1", i+1)
		wg.Add(1)
		go func() {
			aggregate1.ProcessEvent(event)
			wg.Done()
		}()
	}

	// Process events for Aggregate2
	for i := 0; i < numEvents; i++ {
		event := fmt.Sprintf("Event %d for Aggregate2", i+1)
		wg.Add(1)
		go func() {
			aggregate2.ProcessEvent(event)
			wg.Done()
		}()
	}

	// Wait for all aggregates to finish processing events
	wg.Wait()

	printf("All events processed.\n")
}
