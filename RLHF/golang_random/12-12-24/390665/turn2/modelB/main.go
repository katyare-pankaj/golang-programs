package main

import (
	"fmt"
	"strings"
	"sync"
)

// Process a single string by uppercasing it
func processString(wg *sync.WaitGroup, input string, results chan<- string) {
	defer wg.Done() // Signal that this Goroutine is done

	// Simulate some processing time
	fmt.Println("Processing:", input)
	processed := strings.ToUpper(input)
	results <- processed // Send the processed result
}

func main() {
	var wg sync.WaitGroup
	results := make(chan string)

	dataset := []string{"hello", "world", "this", "is", "a", "test"}

	// Start Goroutines to process each string
	for _, item := range dataset {
		wg.Add(1) // Increment the WaitGroup counter
		go processString(&wg, item, results)
	}

	// Close the results channel once all Goroutines are done
	go func() {
		wg.Wait() // Wait for all Goroutines to finish
		close(results)
	}()

	// Collect and print the processed results
	for result := range results {
		fmt.Println("Processed:", result)
	}
}
