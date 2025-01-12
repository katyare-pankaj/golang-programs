package main

import (
	"fmt"
	"sync"
	"time"
)

func makeAPICall(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	// Simulate the API call taking time
	time.Sleep(time.Duration(id) * time.Second)

	fmt.Printf("API call %d completed\n", id)
}

func main() {
	var wg sync.WaitGroup

	// Number of concurrent API calls
	numCalls := 5

	// Add the number of worker goroutines to the WaitGroup
	wg.Add(numCalls)

	// Start the concurrent API calls
	for i := 1; i <= numCalls; i++ {
		go makeAPICall(&wg, i)
	}

	// Wait for all API calls to complete
	wg.Wait()

	fmt.Println("All API calls completed")
}
