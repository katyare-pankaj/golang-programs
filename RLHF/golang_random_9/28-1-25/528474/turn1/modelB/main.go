package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// Worker is a function that processes a task.
type Worker func()

func main() {
	// Number of worker goroutines
	const numWorkers = 100

	// Create a channel to receive tasks
	tasks := make(chan Worker)

	// Create a WaitGroup to wait for all workers to finish
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(&wg, tasks)
	}

	// Serve HTTP requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Create a new task for the worker
		task := func() {
			// Simulate work by sleeping for a random duration
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		}

		// Send the task to the worker channel
		tasks <- task
	})

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	// Wait for all workers to finish their tasks
	wg.Wait()
}

func worker(wg *sync.WaitGroup, tasks <-chan Worker) {
	defer wg.Done()
	for task := range tasks {
		// Execute the task
		task()
	}
}
