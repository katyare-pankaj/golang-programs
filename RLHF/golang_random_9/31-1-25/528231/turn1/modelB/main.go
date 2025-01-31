package main

import (
	"fmt"
	"sync"
)

// Define a callback function type
type DataProcessor func(int) int

// ProcessData reads data from the input channel, applies the callback function,
// and writes the result to the output channel. It runs concurrently using goroutines.
func ProcessData(input <-chan int, output chan<- int, processor DataProcessor, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range input {
		result := processor(data)
		output <- result
	}
}

func main() {
	// Define the input and output channels
	input := make(chan int)
	output := make(chan int)

	// Number of worker goroutines
	numWorkers := 4

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		go ProcessData(input, output, double, &wg)
	}

	// Generate some sample data to process
	go func() {
		defer close(input)
		for i := 1; i <= 10; i++ {
			input <- i
		}
	}()

	// Print the results from the output channel
	go func() {
		for result := range output {
			fmt.Println("Result:", result)
		}
	}()

	// Wait for all processing to complete
	wg.Wait()
	close(output)
	fmt.Println("Processing completed.")
}

// Sample callback function: doubles the input data
func double(data int) int {
	return data * 2
}
