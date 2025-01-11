package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Mock data retrieval function for a single source
func fetchData(ctx context.Context, source string, location string, results chan<- map[string]int, errors chan<- error) {
	var data map[string]int

	// Simulate fetching data with a delay and possible error
	select {
	case <-ctx.Done():
		errors <- ctx.Err()
		return
	case <-time.After(time.Duration(rand.Intn(1000)) * time.Millisecond):
		data = map[string]int{location: sourceValue[source][location]}
	default:
		// Simulate an error
		errors <- fmt.Errorf("error fetching data from %s for %s", source, location)
		return
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
	finalResults := make(map[string]map[string]int)
	for result := range results {
		for location, value := range result {
			if _, ok := finalResults[location]; !ok {
				finalResults[location] = make(map[string]int)
			}
			finalResults[location][value] = value
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
			fmt.Printf("%s: %v\n", location, result)
		}
	}
}
