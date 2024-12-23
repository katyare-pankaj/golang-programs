package main

import (
	"fmt"
	"sync"
)

type itineraryTask func(string) string
type travelItinerary []itineraryTask

func planTrip(destination string, tasks travelItinerary, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, task := range tasks {
		result := task(destination)
		fmt.Printf("Result for %s: %s\n", destination, result)
	}
}

func main() {
	var wg sync.WaitGroup
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

	// Create a slice of destinations
	destinations := []string{"Paris", "Rome", "New York", "Tokyo", "Sydney"}

	// Iterate over the destinations and create a goroutine for each destination
	for _, destination := range destinations {
		wg.Add(1)
		go planTrip(destination, sampleTasks, &wg)
	}

	// Wait for all goroutines to complete their work
	wg.Wait()
	fmt.Println("All trips planned successfully!")
}
