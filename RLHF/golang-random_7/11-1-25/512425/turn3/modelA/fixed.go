package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Simulated data sources
var (
	source1 = map[string]int{"Place A": 100, "Place B": 200}
	source2 = map[string]int{"Place A": 90, "Place B": 210}
	source3 = map[string]int{"Place C": 300}
)

// Simulated function to retrieve data from a single source
func fetchData(ctx context.Context, source string, location string, results chan<- map[string]int, errors chan<- error) {
	var data map[string]int

	// Simulate delay for each source
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	// Simulate error with a probability of 10%
	if rand.Intn(10) == 0 {
		errors <- fmt.Errorf("error fetching data from %s for %s", source, location)
		return
	}

	// Simulate data retrieval
	switch source {
	case "source1":
		data = map[string]int{location: source1[location]}
	case "source2":
		data = map[string]int{location: source2[location]}
	case "source3":
		data = map[string]int{location: source3[location]}
	}

	results <- data
}

func main() {
	locations := []string{"Place A", "Place B", "Place C"}
	sources := []string{"source1", "source2", "source3"}

	// Create channels for results and errors
	results := make(chan map[string]int, len(locations)*len(sources))
	errors := make(chan error)

	var wg sync.WaitGroup

	// Initialize context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Launch goroutines for data retrieval
	for _, location := range locations {
		for _, source := range sources {
			wg.Add(1)
			go func(loc, src string) {
				defer wg.Done()
				fetchData(ctx, src, loc, results, errors)
			}(location, source)
		}
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(results)
	close(errors)

	// Aggregate results and handle errors
	finalResults := make(map[string]int)
	for result := range results {
		for location, value := range result {
			finalResults[location] = value // Store value as the map value
		}
	}

	// Check for any errors
	var hasError bool
	for err := range errors {
		if err != nil {
			fmt.Println("Error:", err)
			hasError = true
		}
	}

	if hasError {
		fmt.Println("Data retrieval failed for some sources.")
	} else {
		// Print final results
		fmt.Println("Final Results:")
		for location, result := range finalResults {
			fmt.Printf("%s: %d\n", location, result) // Print value directly
		}
	}
}
