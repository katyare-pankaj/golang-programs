package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// goroutineFunction represents the function to be executed by goroutines.
func goroutineFunction(id int, wg *sync.WaitGroup, result chan<- int) {
	defer wg.Done() // Decrement the WaitGroup counter when this goroutine exits.

	// Simulate work with a random sleep duration.
	sleepDuration := time.Duration(rand.Intn(500)) * time.Millisecond
	time.Sleep(sleepDuration)

	fmt.Printf("Goroutine %d finished after %v\n", id, sleepDuration)
	// Send the result to the channel for later consumption.
	result <- id * 2
}

func main() {
	const numGoroutines = 5
	var wg sync.WaitGroup

	// Create a buffered channel to store goroutine results.
	result := make(chan int, numGoroutines)

	// Initialize the WaitGroup.
	wg.Add(numGoroutines)

	// Start the goroutines.
	for i := 0; i < numGoroutines; i++ {
		go goroutineFunction(i, &wg, result)
	}

	// All goroutines have been started, now wait for them to finish.
	wg.Wait()

	// Close the result channel to indicate no more values will be sent.
	close(result)

	// Consume the results from the channel.
	totalResult := 0
	for res := range result {
		totalResult += res
	}

	fmt.Println("\nResults from all goroutines:")
	fmt.Println("Total sum:", totalResult)
}
