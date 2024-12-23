package main

import (
	"fmt"
	"sync"
	"time"
)

// Define a function for processing each subtask
func processSubtask(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal to the WaitGroup that this subtask is complete

	// Simulate some work being done
	time.Sleep(time.Duration(id+1) * time.Second)
	fmt.Printf("Subtask %d completed.\n", id)
}

func main() {
	// Create a WaitGroup to track the number of active subtasks
	var wg sync.WaitGroup

	// Define the number of subtasks
	const numSubtasks = 5

	// Start each subtask in a new goroutine
	for i := 0; i < numSubtasks; i++ {
		wg.Add(1)                 // Increment the WaitGroup counter
		go processSubtask(i, &wg) // Start the subtask
	}

	// Wait for all subtasks to complete before proceeding
	wg.Wait()

	fmt.Println("All subtasks have completed. Moving on to the next step...")
}
