package main

import (
	"fmt"
	"sync"
	"time"
)

// Define a contract for data collector
type DataCollector interface {
	CollectData() ([]int, error)
}

// A struct to represent a simple data source
type DataSource struct {
	sourceID int
	data     []int
}

// Implement the DataCollector interface for DataSource
func (ds *DataSource) CollectData() ([]int, error) {
	// Precondition: sourceID must be positive
	if ds.sourceID <= 0 {
		return nil, fmt.Errorf("source ID must be positive")
	}

	// Simulate data collection process
	time.Sleep(time.Duration(ds.sourceID) * time.Millisecond)
	return ds.data, nil
}

// Main function to collect data from multiple sources
func main() {
	var sources []*DataSource
	sources = append(sources, &DataSource{sourceID: 1, data: []int{1, 2, 3}})
	sources = append(sources, &DataSource{sourceID: 2, data: []int{4, 5}})
	sources = append(sources, &DataSource{sourceID: 3, data: []int{6, 7, 8, 9}})

	// Define a WaitGroup to manage goroutines
	var wg sync.WaitGroup
	var results [][]int
	var errors []error

	// Collect data concurrently
	for _, source := range sources {
		wg.Add(1)

		go func(src *DataSource) {
			defer wg.Done()
			data, err := src.CollectData()
			if err != nil {
				errors = append(errors, err)
				return
			}
			results = append(results, data)
		}(source)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Output collected results
	if len(errors) > 0 {
		fmt.Println("Errors occurred during data collection:")
		for _, err := range errors {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Collected data:")
		for _, data := range results {
			fmt.Println(data)
		}
	}
}
