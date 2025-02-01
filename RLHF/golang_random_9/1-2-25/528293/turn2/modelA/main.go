package main

import (
	"fmt"
	"sync"
)

// Define a callback function type
type DataProcessor func(int) int

// ProcessData reads data from the input channel, applies the callback function,
// and writes the result to the corresponding output channel. It runs concurrently using goroutines.
func ProcessData(input <-chan int, output chan<- int, processor DataProcessor, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range input {
		result := processor(data)
		output <- result
	}
}

func main() {
	// Define input channel
	input := make(chan int)

	// Define output channels for each transformation
	outputs := []chan int{
		make(chan int), // For doubling
		make(chan int), // For squaring
	}

	// Define transformation functions
	transformations := []DataProcessor{
		double,
		square,
	}

	// Number of worker goroutines
	numWorkers := len(transformations)

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Start worker goroutines, each responsible for a different transformation
	for i := 0; i < numWorkers; i++ {
		go ProcessData(input, outputs[i], transformations[i], &wg)
	}

	// Generate some sample data to process
	go func() {
		defer close(input)
		for i := 1; i <= 10; i++ {
			input <- i
		}
	}()

	// Start a goroutine to print results from each output channel
	for i := 0; i < numWorkers; i++ {
		go func(i int) {
			for result := range outputs[i] {
				fmt.Printf("Transformation %d result: %d\n", i, result)
			}
		}(i)
	}

	// Wait for all processing to complete
	wg.Wait()

	// Close all output channels after the workers are done
	for _, output := range outputs {
		close(output)
	}

	fmt.Println("Processing completed.")
}

// Sample callback function: doubles the input data
func double(data int) int {
	return data * 2
}

// Sample callback function: squares the input data
func square(data int) int {
	return data * data
}
