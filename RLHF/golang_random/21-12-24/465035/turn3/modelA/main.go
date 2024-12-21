package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulated task processing function
func processTask(task int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that the Goroutine is done
	fmt.Printf("Processing task: %d\n", task)
	time.Sleep(time.Second * 1) // Simulate task duration
	fmt.Printf("Task %d processing completed.\n", task)
}

func main() {
	var wg sync.WaitGroup // Initialize WaitGroup
	tasks := []int{1, 2, 3, 4, 5}

	// Start Goroutines for each task
	for _, task := range tasks {
		wg.Add(1) // Increment the waitgroup for each task
		go processTask(task, &wg)
	}

	// Main thread waits for all tasks to complete
	fmt.Println("Main thread waiting for all tasks to finish...")
	wg.Wait()

	fmt.Println("All tasks completed. Main thread exiting.")
}
