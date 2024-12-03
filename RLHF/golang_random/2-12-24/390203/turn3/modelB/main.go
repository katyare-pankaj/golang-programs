package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var counter int

	// Increment function to safely update shared counter
	increment := func(wg *sync.WaitGroup) {
		defer wg.Done()
		mu.Lock()   // Acquire lock
		counter++   // Critical section
		mu.Unlock() // Release lock
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter) // Expect 10
}
