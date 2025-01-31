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
func work(wg *sync.WaitGroup) {
	defer wg.Done()
	// Add some sleep time to simulate work
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
}

func main() {
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())

	// Start the timer
	start := time.Now()

	// Spawn all the tasks in parallel
	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go work(&wg)
	}

	// Wait for all tasks to complete
	wg.Wait()

	// Stop the timer
	elapsed := time.Since(start)

	// Print the elapsed time
	fmt.Printf("Execution time: %s\n", elapsed)

	//TODO: Add analysis of the Go scheduler's impact on task completion here

}
