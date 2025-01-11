package main

import (
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

// function to retrieve data from each source with concurrent execution
func retrieveData(location string, source string, results chan<- map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()
	var data map[string]int
	switch source {
	case "cache":
		data = map[string]int{location: cache[location]}
	case "db":
		data = map[string]int{location: db[location]}
	case "secondaryDb":
		data = map[string]int{location: secondaryDb[location]}
	}
	// Simulate delay for different sources
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	// Send the result back through the channel
	results <- data
}

func main() {
	locations := []string{"Place A", "Place B", "Place C"}
	sources := []string{"cache", "db", "secondaryDb"}

	wg := sync.WaitGroup{}
	wg.Add(len(locations) * len(sources))

	results := make(chan map[string]int, len(locations)*len(sources))

	// launch goroutines to retrieve data from each source
	for _, location := range locations {
		for _, source := range sources {
			go retrieveData(location, source, results, &wg)
		}
	}

	// wait for all goroutines to finish
	wg.Wait()
	close(results)

	// aggregate results
	finalResults := make(map[string]map[string]int)
	for result := range results {
		for location, value := range result {
			if _, ok := finalResults[location]; !ok {
				finalResults[location] = make(map[string]int)
			}
			finalResults[location]["value"] = value
		}
	}

	// Check for any errors in the channel
	var hasError bool
	for err := range results {
		if err != nil {
			fmt.Println("Error:", err)
			hasError = true
		}
	}

	if hasError {
		fmt.Println("Data retrieval failed for some sources.")
	} else {
		//print final result
		fmt.Println("Final Results:")
		for location, result := range finalResults {
			fmt.Printf("%s: %v\n", location, result)
		}
	}
}
