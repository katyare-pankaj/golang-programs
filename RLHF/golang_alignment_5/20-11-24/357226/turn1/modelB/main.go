package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Third-party library function that might take some time to complete
func thirdPartyLibraryFunction(ctx context.Context, input string) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err() // Return error if context is canceled
	default:
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second) // Simulate variable execution time
		return input + " processed", nil
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Input data for processing
	inputData := []string{"data1", "data2", "data3", "data4", "data5"}

	// Create a channel to receive results
	results := make(chan string)

	// Create a wait group to synchronize Goroutines
	var wg sync.WaitGroup

	// Context for cancellation (optional)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start Goroutines to process data concurrently
	wg.Add(len(inputData))
	for _, data := range inputData {
		go func(data string) {
			defer wg.Done()
			result, err := thirdPartyLibraryFunction(ctx, data)
			if err != nil {
				fmt.Printf("Error processing %s: %v\n", data, err)
				return
			}
			results <- result
		}(data)
	}

	// Wait for all Goroutines to complete
	go func() {
		wg.Wait()
		close(results)
	}()

	// Process results from the channel
	for result := range results {
		fmt.Println("Received result:", result)
	}

	fmt.Println("All processing completed.")
}
