package main

import (
	"fmt"
	"sync"
	"time"
)

// Function to simulate a task
func processTask(taskId int, wg *sync.WaitGroup) {
	defer wg.Done() // Call this when the task is complete to decrement the WaitGroup counter
	fmt.Printf("Starting task %d\n", taskId)
	// Simulate the task duration
	time.Sleep(time.Duration(100+taskId) * time.Millisecond)
	fmt.Printf("Task %d completed\n", taskId)
}

func main() {
	var wg sync.WaitGroup
	numTasks := 5

	// Initialize the WaitGroup with the number of tasks
	wg.Add(numTasks)

	// Start each task in a goroutine
	for i := 1; i <= numTasks; i++ {
		go processTask(i, &wg)
	}

	// Wait for all tasks to finish
	wg.Wait()

	fmt.Println("All tasks completed. Exiting program.")
}
