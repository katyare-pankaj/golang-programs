package main

import (
	"fmt"
	"sync"
	"time"
)

func simulateWork(wg *sync.WaitGroup, id int, duration time.Duration) {
	defer wg.Done()      // Always remember to call Done() when the task is complete
	time.Sleep(duration) // Simulate work by sleeping
	fmt.Printf("Goroutine %d finished its work.\n", id)
}

func main() {
	var wg sync.WaitGroup

	// Simulate starting 5 goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1) // Increment the counter for each goroutine
		go simulateWork(&wg, i, time.Random(1*time.Second, 3*time.Second))
	}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("All goroutines have finished their work.")
}
