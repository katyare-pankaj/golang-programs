package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulate a task that includes opening and closing a resource.
func processTaskWithDefer(id, taskID int) {
	fmt.Printf("Worker %d: Starting Task %d\n", id, taskID)
	// Simulate a resource that needs to be closed
	resource := fmt.Sprintf("Resource-%d.%d", id, taskID)

	// Include defer inside loop -- less efficient
	defer func() {
		closeResource(resource)
	}()

	// Simulate processing time
	time.Sleep(100 * time.Millisecond)

	fmt.Printf("Worker %d: Finished Task %d\n", id, taskID)
}

// Optimized version without repeated defers
func processTaskOptimized(id, taskID int) {
	fmt.Printf("Optimized Worker %d: Starting Task %d\n", id, taskID)
	resource := fmt.Sprintf("Resource-%d.%d", id, taskID)

	// Manually close the resource after use
	closeResource(resource)
	time.Sleep(100 * time.Millisecond)

	fmt.Printf("Optimized Worker %d: Finished Task %d\n", id, taskID)
}

func closeResource(resource string) {
	// Simulate closing the resource
	fmt.Printf("Closing %s\n", resource)
}

func worker(id int, wg *sync.WaitGroup, useDefer bool) {
	defer wg.Done() // This defer is okay as it's not in a loop

	for taskID := 0; taskID < 5; taskID++ {
		if useDefer {
			processTaskWithDefer(id, taskID)
		} else {
			processTaskOptimized(id, taskID)
		}
	}
}

func main() {
	const numWorkers = 3

	fmt.Println("Running with defer:")
	{
		var wg sync.WaitGroup
		wg.Add(numWorkers)
		for i := 0; i < numWorkers; i++ {
			go worker(i, &wg, true) // Use defer
		}
		wg.Wait()
	}

	fmt.Println("\nRunning optimized without defer in loop:")
	{
		var wg sync.WaitGroup
		wg.Add(numWorkers)
		for i := 0; i < numWorkers; i++ {
			go worker(i, &wg, false) // Manual cleanup
		}
		wg.Wait()
	}
}
