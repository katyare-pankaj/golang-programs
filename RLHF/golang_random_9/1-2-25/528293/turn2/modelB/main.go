package main

import (
	"fmt"
	"sync"
)

// Define a struct to hold the configuration for each worker
type WorkerConfig struct {
	Input     <-chan int
	Output    chan<- int
	Processor DataProcessor
}

// Define a callback function type
type DataProcessor func(int) int

// ProcessData reads data from the input channels in WorkerConfigs,
// applies the corresponding callback function, and writes the result to the output channels.
// It runs concurrently using goroutines.
func ProcessData(workers []WorkerConfig, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, worker := range workers {
		go func(worker WorkerConfig) {
			for data := range worker.Input {
				result := worker.Processor(data)
				worker.Output <- result
			}
			close(worker.Output)
		}(worker)
	}
}

func main() {
	// Define the input and output channels for each worker
	input1 := make(chan int)
	output1 := make(chan int)
	input2 := make(chan int)
	output2 := make(chan int)

	// Number of worker goroutines
	numWorkers := 2

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Create a slice of WorkerConfig structs to hold the worker configurations
	workers := []WorkerConfig{
		{
			Input:     input1,
			Output:    output1,
			Processor: double,
		},
		{
			Input:     input2,
			Output:    output2,
			Processor: triple,
		},
	}

	// Start worker goroutines
	go ProcessData(workers, &wg)

	// Generate some sample data to process for each worker
	go func() {
		defer close(input1)
		for i := 1; i <= 5; i++ {
			input1 <- i
		}
	}()

	go func() {
		defer close(input2)
		for i := 1; i <= 5; i++ {
			input2 <- i * 10
		}
	}()

	// Print the results from the output channels
	go func() {
		for result := range output1 {
			fmt.Println("Result from worker 1:", result)
		}
	}()

	go func() {
		for result := range output2 {
			fmt.Println("Result from worker 2:", result)
		}
	}()

	// Wait for all processing to complete
	wg.Wait()
	fmt.Println("Processing completed.")
}

// Sample callback functions: double and triple the input data
func double(data int) int {
	return data * 2
}

func triple(data int) int {
	return data * 3
}
