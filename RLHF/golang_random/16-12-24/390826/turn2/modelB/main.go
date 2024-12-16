package main

import (
	"fmt"
	"sync"
	"time"
)

func incrementCounter(count *int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter when the goroutine is done
	for i := 0; i < 10000; i++ {
		*count++                    // Increment the shared counter
		time.Sleep(time.Nanosecond) // Simulate work
	}
}

func main() {
	var counter int
	var wg sync.WaitGroup // Create a WaitGroup

	// Start multiple goroutines to increment the counter
	wg.Add(5) // Add 5 to the WaitGroup counter
	for i := 0; i < 5; i++ {
		go incrementCounter(&counter, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait() // Wait for the WaitGroup counter to reach 0

	fmt.Println("Final counter value:", counter)
}
