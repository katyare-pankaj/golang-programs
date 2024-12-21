package main

import (
	"fmt"
	"time"
)

func worker(id int, done chan bool) {
	// Simulate work
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Printf("Worker %d is done\n", id)

	// Signal that we are done
	done <- true
}

func main() {
	// Create a channel to signal when workers are done
	done := make(chan bool, 3) // Buffer size of 3 to accommodate 3 workers

	// Start 3 workers
	for i := 0; i < 3; i++ {
		go worker(i+1, done)
	}

	// Wait for all workers to finish
	for i := 0; i < 3; i++ {
		<-done
	}

	fmt.Println("All workers are done")
}
