package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type DataSource struct {
	Name    string
	GetData func(string) (int, error)
}

func retrieveData(location string, source DataSource, results chan<- map[string]int, errors chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	data, err := source.GetData(location)
	if err != nil {
		errors <- fmt.Errorf("error fetching data from %s for %s: %w", source.Name, location, err)
		return
	}
	results <- map[string]int{location: data}
}

func main() {
	locations := []string{"Place A", "Place B", "Place C"}
	dataSources := []DataSource{
		{
			Name: "cache",
			GetData: func(location string) (int, error) {
				// Simulate delay for different sources
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
				// Return cache data or simulate error
				if location == "Place A" {
					return 100, nil
				}
				return 0, errors.New("cache not found")
			},
		},
		{
			Name: "db",
			GetData: func(location string) (int, error) {
				// Simulate delay for different sources
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
				// Return db data or simulate error
				if location == "Place B" {
					return 200, nil
				}
				return 0, errors.New("db not found")
			},
		},
		{
			Name: "secondaryDb",
			GetData: func(location string) (int, error) {
				// Simulate delay for different sources
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
				// Return secondaryDb data or simulate error
				if location == "Place C" {
					return 300, nil
				}
				return 0, errors.New("secondaryDb not found")
			},
		},
	}

	wg := sync.WaitGroup{}
	wg.Add(len(locations) * len(dataSources))

	results := make(chan map[string]int, len(locations)*len(dataSources))
	errors := make(chan error)

	// Launch goroutines to retrieve data from each source
	for _, location := range locations {
		for _, source := range dataSources {
			go retrieveData(location, source, results, errors, &wg)
		}
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(results)
	close(errors)

	// Aggregate results
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
	for err := range errors {
		if err != nil {
			fmt.Println("Error:", err)
			hasError = true
		}
	}

	if hasError {
		fmt.Println("Data retrieval failed for some sources.")
	} else {
		// Print final result
		fmt.Println("Final Results:")
		for location, result := range finalResults {
			fmt.Printf("%s: %v\n", location, result)
		}
	}
}
