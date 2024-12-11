package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var counter int64 // Declare a shared integer

	const numGoroutines = 1000
	var wg sync.WaitGroup

	// Start multiple goroutines that increment the counter
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			// Atomically increment the counter
			atomic.AddInt64(&counter, 1)
			// Simulate some work
			time.Sleep(time.Millisecond * 10)
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Print the final count
	fmt.Println("Final counter value:", counter)
}
