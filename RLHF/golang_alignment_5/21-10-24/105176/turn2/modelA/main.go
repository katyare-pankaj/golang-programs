package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// worker represents a goroutine that processes tasks
func worker(id int, wg *sync.WaitGroup, tasks <-chan int, results chan<- int) {
	defer wg.Done()
	for task := range tasks {
		// Simulate work
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		result := task * task
		results <- result
	}
}

func main() {
	const numWorkers = 5
	const numTasks = 20

	// Create channels
	tasks := make(chan int, numTasks)
	results := make(chan int, numTasks)

	// Create a wait group to wait for worker goroutines to finish
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		go worker(i, &wg, tasks, results)
	}

	// Generate tasks and send them to the workers
	for i := 1; i <= numTasks; i++ {
		tasks <- i
	}

	// Close the tasks channel to signal workers to stop
	close(tasks)

	// Collect results and print them
	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("Result:", result)
	}
}
