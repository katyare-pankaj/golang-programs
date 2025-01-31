package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numTasks = 100000 // You can adjust this number to test with different number of tasks
)

// Simulate computationally intensive task
func work(wg *sync.WaitGroup, times chan<- time.Duration) {
	defer wg.Done()
	// Start the timer for this individual task
	taskStart := time.Now()

	// Add some sleep time to simulate work
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	// Stop the timer and send the elapsed time to the channel
	taskElapsed := time.Since(taskStart)
	times <- taskElapsed
}

func main() {
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())

	// Create a channel to collect times
	times := make(chan time.Duration, numTasks)

	// Start the timer for the whole process
	start := time.Now()

	// Spawn all the tasks in parallel
	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go work(&wg, times)
	}

	// Wait for all tasks to complete
	wg.Wait()
	close(times)

	// Stop the timer for the whole process
	totalElapsed := time.Since(start)

	// Calculate the sum of execution times
	var sum time.Duration
	for t := range times {
		sum += t
	}

	// Calculate the average execution time
	averageTime := sum / numTasks

	// Print the results
	fmt.Printf("Total Execution time: %s\n", totalElapsed)
	fmt.Printf("Average Execution time of individual tasks: %s\n", averageTime)
}
