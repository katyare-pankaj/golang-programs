package main

import (
	"fmt"
	"sync"
	"time"
)

// Function to simulate processing each itinerary task
func processTask(taskID string, duration time.Duration) {
	fmt.Printf("Processing task: %s for %s\n", taskID, duration)
	time.Sleep(duration) // Simulate work by sleeping for a duration
}

// Function to plan a trip, processing all tasks in the itinerary
func planTrip(itinerary [](string, time.Duration), wg *sync.WaitGroup) {
	defer wg.Done() // Decrement counter when the task is done

	for _, task := range itinerary {
		taskID, duration := task
		go processTask(taskID, duration)
	}
}

func main() {
	var wg sync.WaitGroup
	
	// Define different itineraries with various tasks
	itineraries := [][] (string, time.Duration) {
		{
			// Example trip to Paris
			{"Research attractions", 2 * time.Second},
			{"Book flights", 3 * time.Second},
			{"Book accommodation", 1 * time.Second},
		},
		{
			// Example trip to Rome
			{"Research attractions", 2 * time.Second},
			{"Order travel insurance", 1 * time.Second},
			{"Plan transportation", 3 * time.Second},
		},
	}
	
	// Iterate over each itinerary and execute concurrently
	for _, itinerary := range itineraries {
		wg.Add(1) // Increment counter before starting goroutine
		go planTrip(itinerary, &wg)
	}
	
	// Wait for all trips to be planned
	wg.Wait()

	fmt.Println("All itineraries processed successfully!")
}