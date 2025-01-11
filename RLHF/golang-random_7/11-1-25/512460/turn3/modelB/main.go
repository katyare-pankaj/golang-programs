package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// Shared resource: A slice of integers that will be accessed by multiple goroutines.
var sharedResource = make([]int, 100)

// Worker function: This function will read and write to the shared resource.
func worker(id int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		// Generate a random index to access the shared resource
		index := rand.Intn(len(sharedResource))

		// Request access to the shared resource by sending an index on the channel
		ch <- index

		// Read and modify the value at the requested index
		sharedResource[index]++

		// Release the shared resource by receiving from the channel
		<-ch
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 1) // Buffered channel with a capacity of 1 to limit resource contention

	// Start multiple workers
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go worker(i, &wg, ch)
	}

	// Wait for all workers to complete
	wg.Wait()

	// Check for race conditions by verifying that the sum of the shared resource is correct
	sum := 0
	for _, value := range sharedResource {
		sum += value
	}
	fmt.Println("Sum of the shared resource:", sum)

	// If the sum is not correct, it indicates a race condition
	if sum != 10000 {
		fmt.Println("Race condition detected!")
	}
}
