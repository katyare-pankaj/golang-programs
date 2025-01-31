package main

import (
	"fmt"
	"sync"
)

// Define a callback function type
type DataProcessor func(int) int

// Map to store dynamic callbacks
var callbacks = map[string]DataProcessor{}
var callbackMutex sync.RWMutex

// ProcessData reads data from the input channel, applies the callback function,
// and writes the result to the corresponding output channel. It runs concurrently using goroutines.
func ProcessData(input <-chan int, output chan<- int, callbackName string, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range input {
		callbackMutex.RLock()
		processor := callbacks[callbackName]
		callbackMutex.RUnlock()
		if processor == nil {
			fmt.Println("Error: No callback function found for", callbackName)
			continue
		}
		result := processor(data)
		output <- result
	}
}

func main() {
	// Define input channel
	input := make(chan int)

	// Define output channels for each transformation (no need for slice now)
	outputDouble := make(chan int)
	outputSquare := make(chan int)

	// Number of worker goroutines
	numWorkers := 2

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Register initial callbacks
	registerCallback("double", double)
	registerCallback("square", square)

	// Start worker goroutines
	go ProcessData(input, outputDouble, "double", &wg)
	go ProcessData(input, outputSquare, "square", &wg)

	// Generate some sample data to process
	go func() {
		defer close(input)
		for i := 1; i <= 10; i++ {
			input <- i
		}
	}()

	// Start a goroutine to print results from each output channel
	go func() {
		for result := range outputDouble {
			fmt.Println("Result:", result)
		}
	}()

	go func() {
		for result := range outputSquare {
			fmt.Println("Result (Squared):", result)
		}
	}()

	// Register a new callback at runtime
	go func() {
		callbacks["cube"] = cube
	}()

	// Wait for all processing to complete
	wg.Wait()
	close(outputDouble)
	close(outputSquare)
	fmt.Println("Processing completed.")
}

// Sample callback functions: doubling and squaring
func double(data int) int {
	return data * 2
}

func square(data int) int {
	return data * data
}

func cube(data int) int {
	return data * data * data
}

func registerCallback(name string, processor DataProcessor) {
	callbackMutex.Lock()
	defer callbackMutex.Unlock()
	callbacks[name] = processor
}
