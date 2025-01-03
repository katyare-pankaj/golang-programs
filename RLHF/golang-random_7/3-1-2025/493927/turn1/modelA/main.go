package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// Define a simple migration task
type MigrationTask struct {
	id   int
	data string
}

// Perform a data migration (simulated for demonstration)
func MigrateData(task MigrationTask, wg *sync.WaitGroup, results chan error) {
	defer wg.Done()

	// Simulate data migration work
	fmt.Printf("Migrating data %d...\n", task.id)
	time.Sleep(2 * time.Second)

	// Simulate an error for 20% of tasks
	if task.id%5 == 0 {
		results <- errors.New("failed to migrate data")
	} else {
		results <- nil
	}
}

func main() {
	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Create a channel to collect results (errors in this case)
	results := make(chan error, 100)

	// Number of tasks
	numTasks := 20

	// Launch goroutines to perform migrations
	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go MigrateData(MigrationTask{id: i + 1, data: fmt.Sprintf("Task %d data", i+1)}, &wg, results)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Close the results channel after processing all tasks
	close(results)

	// Handle results
	fmt.Println("Migration results:")
	for result := range results {
		if result != nil {
			fmt.Printf("Error: %v\n", result)
		} else {
			fmt.Println("Migration successful.")
		}
	}
}
