package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numTasks = 1000
)

var (
	taskTimes []time.Duration // Global slice to store task times
	mutex     sync.Mutex      // Mutex to lock access to taskTimes slice
)

func work(wg *sync.WaitGroup) {
	defer wg.Done()

	// Record the start time of the task
	start := time.Now()

	// Simulate computationally intensive task
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	// Record the end time of the task
	elapsed := time.Since(start)

	// Acquire lock to safely append to the taskTimes slice
	mutex.Lock()
	taskTimes = append(taskTimes, elapsed)
	mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())

	// Start the timer
	start := time.Now()

	// Initialize taskTimes slice
	taskTimes = make([]time.Duration, 0, numTasks)

	// Spawn all the tasks in parallel
	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go work(&wg)
	}

	wg.Wait() // Wait for all tasks to complete

	// Stop the timer for the overall execution
	overallElapsed := time.Since(start)

	// Calculate the average time
	totalTime := time.Duration(0)
	for _, t := range taskTimes {
		totalTime += t
	}
	averageTime := totalTime / time.Duration(len(taskTimes))

	// Print results
	fmt.Printf("Average task execution time: %s\n", averageTime)
	fmt.Printf("Overall execution time: %s\n", overallElapsed)
	fmt.Println("Task times:")
	for _, t := range taskTimes {
		fmt.Printf("  %s\n", t)
	}
}
