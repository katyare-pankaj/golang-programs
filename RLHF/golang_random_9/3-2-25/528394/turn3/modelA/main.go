package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// Shared resource using atomic counter
	counter int64
)

func incrementWithAtomic(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		// Use atomic add to increment the counter safely
		atomic.AddInt64(&counter, 1)
		currentValue := atomic.LoadInt64(&counter) // Ensure atomic read
		fmt.Printf("With Atomic: Goroutine %d incremented counter to %d\n", id, currentValue)
		time.Sleep(time.Millisecond) // Simulate some processing
	}
}

func main() {
	var wg sync.WaitGroup
	numGoroutines := 5

	fmt.Println("Incrementing counter with Atomic Operations:")
	counter = 0
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go incrementWithAtomic(i, &wg)
	}
	wg.Wait()

	fmt.Printf("Final counter value with Atomic Operations: %d\n", counter)
}
