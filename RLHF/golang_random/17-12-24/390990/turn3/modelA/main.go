package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	var wg sync.WaitGroup

	incrementCounter := func() {
		for i := 0; i < 10000; i++ {
			// Atomically increment the counter
			atomic.AddInt64(&counter, 1)
		}
		wg.Done()
	}

	// Launch multiple goroutines to increment the counter
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go incrementCounter()
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Final counter value:", counter)
}
