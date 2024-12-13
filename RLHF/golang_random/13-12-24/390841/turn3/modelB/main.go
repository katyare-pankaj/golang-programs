package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// trackTask executes a task and measures its execution time
func trackTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	start := time.Now()

	// Simulate work with random sleep
	duration := time.Duration(rand.Intn(100)) * time.Millisecond
	time.Sleep(duration)

	elapsed := time.Since(start)
	fmt.Printf("Task %d completed in %v\n", id, elapsed)
}

func main() {
	const numTasks = 10
	var wg sync.WaitGroup

	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go trackTask(i, &wg)
	}

	wg.Wait() // Wait for all tasks to finish
}
