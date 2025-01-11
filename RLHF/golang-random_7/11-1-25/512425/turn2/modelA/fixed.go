package main

import (
	"errors" // Importing the errors package
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Mock data sources
var (
	cache       = map[string]int{"Place A": 100, "Place B": 200}
	db          = map[string]int{"Place A": 90, "Place B": 210}
	secondaryDb = map[string]int{"Place C": 300}
)

// function to retrieve data from a single source
func retrieveDataFromSource(location string, source string, results chan<- map[string]int, errChan chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	var data map[string]int
	var err error
	switch source {
	case "cache":
		data = map[string]int{location: cache[location]}
	case "db":
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // Simulate delay
		if rand.Intn(2) == 0 {
			err = errors.New("Database error")
		} else {
			data = map[string]int{location: db[location]}
		}
	case "secondaryDb":
		data = map[string]int{location: secondaryDb[location]}
	default:
		err = errors.New("Invalid source")
	}

	if err != nil {
		errChan <- err // Send error to the channel
		return
	}

	results <- data // Send results to the channel
}

func main() {
	locations := []string{"Place A", "Place B", "Place C"}
	sources := []string{"cache", "db", "secondaryDb"}

	wg := sync.WaitGroup{}
	wg.Add(len(locations) * len(sources))

	results := make(chan map[string]int, len(locations)*len(sources))
	errChan := make(chan error, len(locations)*len(sources)) // Renamed to errChan to avoid conflict

	// Launch goroutines to retrieve data from each source
	for _, location := range locations {
		for _, source := range sources {
			go retrieveDataFromSource(location, source, results, errChan, &wg)
		}
	}

	// Close the results and errors channels when all goroutines are finished
	go func() {
		wg.Wait()
		close(results)
		close(errChan)
	}()

	// Aggregate results and handle errors
	finalResults := make(map[string]map[string]int)
	var hasError bool
	for result := range results {
		for location, value := range result {
			if _, ok := finalResults[location]; !ok {
				finalResults[location] = make(map[string]int)
			}
			// Instead of using "value" under "value", just assign the actual value
			finalResults[location] = map[string]int{"value": value}
		}
	}

	// Check for errors and report
	for err := range errChan {
		if err != nil {
			fmt.Println("Error:", err)
			hasError = true
		}
	}

	if hasError {
		fmt.Println("Data retrieval failed for some sources.")
	} else {
		// Print the final result
		fmt.Println("Final Results:")
		for location, result := range finalResults {
			fmt.Printf("%s: %v\n", location, result)
		}
	}
}
