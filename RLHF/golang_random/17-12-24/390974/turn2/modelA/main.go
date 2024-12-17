package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Data represents the type of data we'll be processing
type Data struct {
	ID int
}

// Process simulates processing a piece of data
func Process(data Data) {
	// Simulate work by sleeping
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	fmt.Printf("Processed Data ID: %d\n", data.ID)
}

func main() {
	// Set up context for cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	const (
		numWorkers   = 5   // Number of worker goroutines
		dataBuffer   = 100 // Size of the buffer for incoming data
		numDataItems = 500 // Total items to process
	)

	// Create a buffered channel for incoming data
	dataChan := make(chan Data, dataBuffer)

	// WaitGroup to manage goroutine completion
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for {
				select {
				case data, ok := <-dataChan:
					if !ok {
						return // Channel closed, exit the goroutine
					}
					Process(data)
				case <-ctx.Done():
					return // Context canceled, exit the goroutine
				}
			}
		}(i)
	}

	// Simulate data production
	go func() {
		for i := 0; i < numDataItems; i++ {
			data := Data{ID: i}
			dataChan <- data // Send data to the channel
		}
		close(dataChan) // Close the channel after sending all data
	}()

	// Wait for all workers to finish processing
	wg.Wait()

	fmt.Println("All data processed.")
}
