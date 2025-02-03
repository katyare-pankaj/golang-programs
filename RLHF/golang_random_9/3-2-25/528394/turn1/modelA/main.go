package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	// Shared resource
	counter int
)

// simulateDatabaseOperation simulates a database operation by incrementing a counter.
func simulateDatabaseOperation(id int, wg *sync.WaitGroup, mtx *sync.Mutex) {
	defer wg.Done()

	for i := 0; i < 100; i++ {
		// Use mutex to lock the shared resource
		mtx.Lock()
		counter++
		fmt.Printf("Goroutine %d: incremented counter to %d\n", id, counter)
		mtx.Unlock()
		time.Sleep(10 * time.Millisecond) // simulate time-consuming operation
	}
}

func main() {
	var wg sync.WaitGroup
	var mtx sync.Mutex
	numGoroutines := 5
	counter = 0

	// Using WaitGroup and Mutex together to manage concurrency
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go simulateDatabaseOperation(i, &wg, &mtx)
	}

	wg.Wait() // waits for all goroutines to complete

	fmt.Printf("Final counter value: %d\n", counter)
}
