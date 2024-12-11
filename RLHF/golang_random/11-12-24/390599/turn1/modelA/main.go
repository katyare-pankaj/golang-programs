package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mutex.Lock()   // Lock access to the counter
		counter++      // Increment the counter
		mutex.Unlock() // Unlock access to the counter
	}
}

func main() {
	var wg sync.WaitGroup

	// Launch 10 goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Printf("Final counter value: %d\n", counter)
}
