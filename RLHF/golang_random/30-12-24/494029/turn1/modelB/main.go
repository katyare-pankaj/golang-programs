package main

import (
	"fmt"
	"sync"
)

func main() {
	// Define an array to store the collected data
	var data []string
	// Create a WaitGroup to keep track of the number of worker goroutines
	var wg sync.WaitGroup

	// Enforce precondition: Input should be a non-nil pointer to a slice
	if data == nil {
		panic("Input data slice is nil")
	}

	// Add one to WaitGroup for each worker goroutine
	wg.Add(3)

	// Launch worker goroutines to collect data from different sources
	go collectDataFromSource1(&data, &wg)
	go collectDataFromSource2(&data, &wg)
	go collectDataFromSource3(&data, &wg)

	// Wait for all worker goroutines to complete
	wg.Wait()

	// Postcondition: All data sources have been collected
	fmt.Println("Data collection completed.")
	fmt.Println("Collected Data:", data)
}

// Function to collect data from source 1
func collectDataFromSource1(data *[]string, wg *sync.WaitGroup) {
	// Defer a call to Done() to signal the WaitGroup that this worker is done
	defer wg.Done()

	// Example data collection logic from source 1
	source1Data := []string{"Data from source 1-1", "Data from source 1-2", "Data from source 1-3"}
	*data = append(*data, source1Data...)

	// Function Invariant: After appending data, the data slice should not be empty
	if len(*data) == 0 {
		panic("Data slice is empty after collecting data from source 1")
	}
}

// Function to collect data from source 2 (follows similar structure to source 1)
func collectDataFromSource2(data *[]string, wg *sync.WaitGroup) {
	defer wg.Done()
	source2Data := []string{"Data from source 2-1", "Data from source 2-2"}
	*data = append(*data, source2Data...)
	if len(*data) == 0 {
		panic("Data slice is empty after collecting data from source 2")
	}
}

// Function to collect data from source 3 (follows similar structure to source 1)
func collectDataFromSource3(data *[]string, wg *sync.WaitGroup) {
	defer wg.Done()
	source3Data := []string{"Data from source 3-1"}
	*data = append(*data, source3Data...)
	if len(*data) == 0 {
		panic("Data slice is empty after collecting data from source 3")
	}
}
