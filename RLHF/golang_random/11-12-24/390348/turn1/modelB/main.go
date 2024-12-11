package main

import (
	"fmt"
	"sync"
	"time"
)

// Data structure to hold aggregated data
type AggregatedData struct {
	Count int
	Sum   float64
}

// Function to simulate processing a data stream
func processStream(stream <-chan int, wg *sync.WaitGroup, agg *AggregatedData) {
	defer wg.Done()
	for data := range stream {
		agg.Count++
		agg.Sum += float64(data)
		time.Sleep(time.Millisecond * 10) // Simulate processing time
	}
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 4
	dataStreams := make([]chan int, numWorkers)
	aggregatedData := make([]AggregatedData, numWorkers)

	// Create data streams
	for i := 0; i < numWorkers; i++ {
		dataStreams[i] = make(chan int)
		go func(stream chan<- int) {
			defer close(stream)
			for j := 0; j < 1000; j++ {
				stream <- i*100 + j // Generate sample data
			}
		}(dataStreams[i])
	}

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go processStream(dataStreams[i], &wg, &aggregatedData[i])
	}

	// Wait for all workers to complete
	wg.Wait()

	// Aggregate results
	totalCount := 0
	totalSum := 0.0
	for _, data := range aggregatedData {
		totalCount += data.Count
		totalSum += data.Sum
	}

	// Output results
	fmt.Printf("Total Count: %d\n", totalCount)
	fmt.Printf("Total Sum: %.2f\n", totalSum)
}
