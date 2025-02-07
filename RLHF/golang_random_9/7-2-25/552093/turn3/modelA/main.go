package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(name string, duration time.Duration, wg *sync.WaitGroup) {
	// Signal that the goroutine has finished executing
	defer wg.Done()

	fmt.Printf("%s starting\n", name)
	time.Sleep(duration) // Simulate work by sleeping
	fmt.Printf("%s finished\n", name)
}

func main() {
	var wg sync.WaitGroup

	// We have two goroutines to wait for
	wg.Add(2)

	go worker("Worker 1", 3*time.Second, &wg) // Task that takes 3 seconds
	go worker("Worker 2", 1*time.Second, &wg) // Task that takes 1 second

	// Wait for both goroutines to finish
	wg.Wait()

	fmt.Println("Both workers have completed their tasks.")
}
