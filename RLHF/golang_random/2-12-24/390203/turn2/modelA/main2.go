package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Start 3 goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the WaitGroup counter

		go func(id int) {
			defer wg.Done() // Decrement the counter when the goroutine completes
			fmt.Printf("Goroutine %d starting\n", id)
			time.Sleep(time.Duration(id) * time.Second) // Simulate work
			fmt.Printf("Goroutine %d done\n", id)
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("All goroutines finished")
}
