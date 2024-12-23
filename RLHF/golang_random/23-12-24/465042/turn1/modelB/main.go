package main

import (
	"fmt"
	"sync"
)

type itineraryTask func(string) string
type travelItinerary []itineraryTask

// This function represents the process of planning a trip to a new destination
func planTrip(destination string, tasks travelItinerary, wg *sync.WaitGroup) {
	// We defer the call to Done() to ensure that the WaitGroup is notified even if the function exits unexpectedly
	defer wg.Done()

	// Process each task in the itinerary concurrently
	for _, task := range tasks {
		go func(t itineraryTask) {
			result := t(destination)
			fmt.Printf("Result for %s: %s\n", destination, result)
		}(task)
	}
}

func main() {
	// Create a WaitGroup to keep track of the number of pending goroutines
	var wg sync.WaitGroup

	// Define some sample itinerary tasks
	sampleTasks := travelItinerary{
		func(destination string) string {
			return fmt.Sprintf("Research attractions in %s", destination)
		},
		func(destination string) string {
			return fmt.Sprintf("Buy flight tickets to %s", destination)
		},
		func(destination string) string {
			return fmt.Sprintf("Book accommodation in %s", destination)
		},
	}

	// Plan trips to different destinations concurrently
	destinations := []string{"Paris", "Rome", "New York"}
	for _, destination := range destinations {
		// Add one to the WaitGroup for each goroutine that will be started
		wg.Add(1)
		// Start a goroutine to plan the trip to the destination
		go planTrip(destination, sampleTasks, &wg)
	}

	// Wait for all goroutines to complete their work
	wg.Wait()

	fmt.Println("All trips planned successfully!")
}
