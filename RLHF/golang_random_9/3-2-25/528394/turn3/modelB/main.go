package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	counter int64 // Use int64 for atomic operations
)

func incrementWithoutMutex(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		// Use atomic.AddInt64 to increment the counter safely
		atomic.AddInt64(&counter, 1)
		fmt.Printf("Without Mutex: Goroutine %d incremented counter to %d\n", id, counter)
		time.Sleep(time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	numGoroutines := 5

	// Part 1: Without Mutex
	fmt.Println("Incrementing counter without Mutex using Atomic Operations:")
	counter = 0
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go incrementWithoutMutex(i, &wg)
	}
	wg.Wait()
	fmt.Printf("Final counter value without Mutex: %d\n", counter)
}
