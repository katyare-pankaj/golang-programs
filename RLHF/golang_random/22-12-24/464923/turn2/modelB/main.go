package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Function to simulate a data processing task
func processData(taskID int, data []int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Task %d: Processing data...\n", taskID)
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond) // Simulate varying task execution times
	fmt.Printf("Task %d: Processing completed. Result: %v\n", taskID, data)
}

func main() {
	// Number of concurrent tasks
	const numTasks = 3
	// Initialize WaitGroup
	var wg sync.WaitGroup
	// Simulate some data to process
	data := []int{1, 2, 3, 4, 5}

	// Start each task in a goroutine and add it to the WaitGroup
	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go processData(i+1, data, &wg)
	}

	// Wait for all tasks to complete before proceeding with the next step
	fmt.Println("Waiting for all tasks to complete...")
	wg.Wait()
	fmt.Println("All tasks have completed.")

	// Next step in the pipeline
	finalResults := []int{}
	for i := 0; i < len(data); i++ {
		finalResults = append(finalResults, data[i]*data[i])
	}
	fmt.Println("Final results:", finalResults)
}
