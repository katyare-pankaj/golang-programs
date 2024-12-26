package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	numGoroutines := 10

	// Set the wait group counter
	wg.Add(numGoroutines)

	// Start goroutines
	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Duration(i) * time.Second)
			fmt.Printf("Goroutine %d finished\n", i)
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	fmt.Println("All goroutines finished.")
}
