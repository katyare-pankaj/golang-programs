package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark the worker as done when the function exits

	// Simulate doing some work with a random duration
	duration := time.Duration(rand.Intn(1000)) * time.Millisecond
	fmt.Printf("Worker %d starting work for %v\n", id, duration)
	time.Sleep(duration) // Simulating work
	fmt.Printf("Worker %d finished work\n", id)
}

func main() {
	var wg sync.WaitGroup // Create a WaitGroup to coordinate goroutines
	numWorkers := 5       // Number of worker goroutines to launch

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)         // Increment the WaitGroup counter for each goroutine
		go worker(i, &wg) // Start a new goroutine
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All workers completed.")
}
