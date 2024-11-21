package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	const numTasks = 100

	// Create a channel to communicate errors
	errorChannel := make(chan error, numTasks)

	// Function to simulate a task that might panic
	func task(id int, errorChannel chan error) {
		defer wg.Done()

		// Simulate some work
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

		// Introduce a potential panic
		if rand.Intn(2) == 0 {
			panic(fmt.Sprintf("Task %d failed", id))
		}
	}

	// Start tasks concurrently
	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go task(i, errorChannel)
	}

	// Wait for all tasks to finish
	wg.Wait()

	// Collect and process errors from the error channel
	for err := range errorChannel {
		if err != nil {
			log.Println("Recovered error:", err)
		}
	}
}