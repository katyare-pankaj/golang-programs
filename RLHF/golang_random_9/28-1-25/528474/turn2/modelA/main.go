package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Simulates a task with variable workload duration.
func performTask(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Indicate that the task is complete.

	// Generate a random duration to simulate work.
	duration := time.Duration(rand.Intn(2000)) * time.Millisecond
	fmt.Printf("Task %d is starting and will take %v\n", id, duration)

	// Simulate workload.
	time.Sleep(duration)

	fmt.Printf("Task %d completed\n", id)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	const numTasks = 10 // Number of tasks

	var wg sync.WaitGroup
	wg.Add(numTasks)

	for i := 0; i < numTasks; i++ {
		go performTask(i, &wg)
	}

	// Wait for all tasks to complete.
	wg.Wait()

	fmt.Println("All tasks have completed.")
}
