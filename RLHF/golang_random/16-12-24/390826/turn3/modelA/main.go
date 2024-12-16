package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal completion when the function exits
	fmt.Printf("Worker %d starting\n", id)
	// Simulate some processing
	// time.Sleep(time.Second) // Optional: simulate work with a sleep
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup // Create a WaitGroup
	const numWorkers = 5

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)         // Increment the WaitGroup counter
		go worker(i, &wg) // Start a new goroutine
	}

	wg.Wait() // Wait for all workers to complete
	fmt.Println("All workers finished")
}
