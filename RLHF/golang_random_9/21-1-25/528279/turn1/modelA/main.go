package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create a channel to receive streaming data
	dataStream := make(chan int)
	var wg sync.WaitGroup

	// Shared slice to accumulate processed results
	var results []int
	var mu sync.Mutex

	// Start a goroutine to produce data
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			dataStream <- i // Simulating stream data
		}
		close(dataStream)
	}()

	// Start several consumer goroutines
	numConsumers := 3
	for i := 0; i < numConsumers; i++ {
		wg.Add(1)
		go func(consumerID int) {
			defer wg.Done()
			for data := range dataStream {
				// Process incoming data
				processed := processData(consumerID, data)

				mu.Lock()
				results = append(results, processed)
				mu.Unlock()
			}
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Print final accumulated results
	fmt.Println("Final results:", results)
}

func processData(consumerID, data int) int {
	// Simulate some processing
	fmt.Printf("Consumer %d processing data %d\n", consumerID, data)
	return data * 2 // Example transformation of the data
}
