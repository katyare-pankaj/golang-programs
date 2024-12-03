package main

import (
	"fmt"
	"sync"
	"time"
)

// Define a type for the data processing function, which returns an error
type DataProcessor func(data string) error

// Higher-order function to process data concurrently
func ProcessConcurrently(data []string, processor DataProcessor) {
	var wg sync.WaitGroup
	errors := make([]error, len(data))

	// Function to process each data item and capture its error
	processItem := func(i int, data string) {
		defer wg.Done()
		errors[i] = processor(data)
	}

	// Start processing each item in a new Goroutine
	for i, item := range data {
		wg.Add(1)
		go processItem(i, item)
	}

	// Wait for all Goroutines to finish
	wg.Wait()

	// Aggregate and display errors
	var aggregatedErrors []error
	for _, err := range errors {
		if err != nil {
			aggregatedErrors = append(aggregatedErrors, err)
		}
	}

	if len(aggregatedErrors) > 0 {
		fmt.Println("Errors encountered:")
		for _, err := range aggregatedErrors {
			fmt.Printf("%v\n", err)
		}
	} else {
		fmt.Println("All data processed successfully.")
	}
}

func main() {
	// Sample data to process
	data := []string{
		"Item 1",
		"Item 2",
		"Item 3",
	}

	// Define a simple data processor function
	processor := func(data string) error {
		fmt.Println("Processing:", data)
		time.Sleep(1 * time.Second) // Simulate processing time

		// Introduce a simulated error for one item
		if data == "Item 2" {
			return fmt.Errorf("Error processing %s", data)
		}

		return nil
	}

	// Process data concurrently
	ProcessConcurrently(data, processor)
}
