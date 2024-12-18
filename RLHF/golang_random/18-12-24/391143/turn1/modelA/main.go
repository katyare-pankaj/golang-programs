package main

import (
	"fmt"
	"time"
)

func worker(id int) {
	for {
		time.Sleep(time.Millisecond * 200) // Simulate work or block
		fmt.Printf("Worker %d doing work...\n", id)
	}
}

func main() {
	// Create a few worker Goroutines
	numWorkers := 5
	for i := 1; i <= numWorkers; i++ {
		go worker(i)
	}

	fmt.Println("Goroutines are running. Press CTRL+C to exit.")
	select {} // Busy-wait until interrupted
}
