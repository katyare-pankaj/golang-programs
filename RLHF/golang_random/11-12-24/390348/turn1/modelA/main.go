package main

import (
	"fmt"
	"sync"
)

// Data stream type
type DataStream struct {
	ID    int
	Value int
}

// Aggregator function
func aggregate(streams <-chan DataStream, wg *sync.WaitGroup, resultChan chan<- int) {
	defer wg.Done()
	total := 0
	for stream := range streams {
		total += stream.Value
	}
	resultChan <- total
}

// Main function
func main() {
	const numStreams = 5
	var wg sync.WaitGroup
	clientCount := 3

	// Channel to receive data streams
	dataStreams := make(chan DataStream, clientCount)
	// Channel to receive the final result
	resultChan := make(chan int, numStreams)

	// Start aggregating goroutines
	for i := 0; i < numStreams; i++ {
		wg.Add(1)
		go aggregate(dataStreams, &wg, resultChan)
	}

	// Simulate sending data to the dataStreams channel
	data := []DataStream{
		{ID: 1, Value: 10},
		{ID: 2, Value: 20},
		{ID: 3, Value: 30},
		{ID: 4, Value: 40},
		{ID: 5, Value: 50},
	}

	for _, d := range data {
		dataStreams <- d
	}
	close(dataStreams) // Close the channel so goroutines can finish

	// Wait for all goroutines to finish
	wg.Wait()
	close(resultChan)

	// Aggregate results from resultChan
	finalTotal := 0
	for result := range resultChan {
		finalTotal += result
	}

	fmt.Printf("Final Total: %d\n", finalTotal)
}
