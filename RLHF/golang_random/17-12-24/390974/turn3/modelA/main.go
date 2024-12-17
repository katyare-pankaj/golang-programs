package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Data represents the type of work
type Data struct {
	ID int
}

// Process simulates data processing
func Process(data Data) {
	// Simulate work
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	fmt.Printf("Processed Data ID: %d\n", data.ID)
}

func main() {
	const (
		numWorkers   = 10  // Number of workers
		dataBuffer   = 100 // Size of the buffer for incoming data channel
		numDataItems = 500 // Total number of data items to process
	)

	// Create context for cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Buffered channel for incoming data
	dataChan := make(chan Data, dataBuffer)

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
						return // Channel closed, exit goroutine
					}
					Process(data)
				case <-ctx.Done():
					return // Context canceled, exit goroutine
				}
			}
		}(i)
	}

	// Simulate data generation
	go func() {
		for i := 0; i < numDataItems; i++ {
			dataChan <- Data{ID: i} // Send data into the channel
		}
		close(dataChan) // Close the channel when done producing
	}()

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All data processed.")
}
