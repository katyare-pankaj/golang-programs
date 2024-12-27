package main

import (
	"fmt"
	"sync"
)

const batchSize = 10

// processTask simulates processing a single task.
func processTask(task string, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simulate task processing (Uncomment if needed for debugging)
	// fmt.Println("Processing task:", task)
}

func main() {
	// Define the list of tasks to process
	tasks := []string{
		"Task1", "Task2", "Task3", "Task4", "Task5", "Task6", "Task7", "Task8", "Task9", "Task10",
		"Task11", "Task12", "Task13", "Task14", "Task15", "Task16", "Task17", "Task18", "Task19", "Task20",
		// Add up to "Task1000" as needed
	}

	var wg sync.WaitGroup

	// Process tasks in batches
	for i := 0; i < len(tasks); i += batchSize {
		end := i + batchSize
		if end > len(tasks) {
			end = len(tasks)
		}

		wg.Add(end - i)

		// Launch a goroutine for each task in the batch
		for j := i; j < end; j++ {
			go processTask(tasks[j], &wg)
		}

		// Wait for the current batch to complete before moving to the next
		wg.Wait()
	}

	fmt.Println("All tasks completed.")
}
