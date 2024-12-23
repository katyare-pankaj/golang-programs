package main

import (
	"fmt"
	"sync"
	"time"
)

// A function that simulates a task taking different input values.
func processTask(id int, input string, wg *sync.WaitGroup) {
	defer wg.Done()                             // Decrement the WaitGroup counter when the task is complete
	time.Sleep(time.Duration(id) * time.Second) // Simulate work taking different times
	fmt.Printf("Task %d with input '%s' completed\n", id, input)
}

func main() {
	var wg sync.WaitGroup

	// Define the input values for the tasks
	inputs := []string{"Input A", "Input B", "Input C"}

	// Start a goroutine for each input value
	for id, input := range inputs {
		wg.Add(1) // Increment the WaitGroup counter for each goroutine
		go func(id int, input string) {
			processTask(id, input, &wg)
		}(id, input)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	fmt.Println("All tasks have completed")
}
