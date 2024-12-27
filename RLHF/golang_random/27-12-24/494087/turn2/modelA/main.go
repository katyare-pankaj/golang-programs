package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done() // Decrement the counter when the goroutine finishes
	fmt.Println("Worker", id, "starting")
	time.Sleep(2 * time.Second)
	fmt.Println("Worker", id, "completed")
}

func main() {
	var wg sync.WaitGroup

	// Start 10 goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1) // Increment the counter
		go worker(&wg, i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All workers completed.")
}
