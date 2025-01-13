package main

import (
	"fmt"
	"sync"
	"time"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simulate some work being done by the task
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Printf("Task %d is complete.\n", id)
}

func main() {
	var wg sync.WaitGroup
	numTasks := 5
	wg.Add(numTasks)
	// Start the tasks concurrently
	for i := 1; i <= numTasks; i++ {
		go task(i, &wg)
	}
	// Wait for all tasks to finish
	fmt.Println("Waiting for all tasks to complete...")
	wg.Wait()
	fmt.Println("All tasks have completed.")
}
