package main

import (
	"fmt"
	"sync"
	"time"
)

// Task represents a generic task structure that will be processed by workers
type Task struct {
	ID      int
	Payload string
}

// worker function performs tasks and sends the result back to the results channel
func worker(id int, tasks <-chan Task, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the waitgroup counter when the worker exits

	for task := range tasks {
		fmt.Printf("Worker %d started task %d\n", id, task.ID)
		// Simulate processing time
		time.Sleep(2 * time.Second)
		result := fmt.Sprintf("Worker %d processed task with payload: %s", id, task.Payload)
		fmt.Printf("Worker %d finished task %d\n", id, task.ID)
		results <- result // Send the result back through the results channel
	}
}

func main() {
	const numWorkers = 3
	const numTasks = 10

	tasks := make(chan Task, numTasks)     // Channel for sending tasks to workers
	results := make(chan string, numTasks) // Channel for receiving results from workers

	var wg sync.WaitGroup

	// Start worker pool
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, tasks, results, &wg)
	}

	// Send tasks to the workers
	go func() {
		for i := 1; i <= numTasks; i++ {
			tasks <- Task{ID: i, Payload: fmt.Sprintf("Task-%d", i)}
		}
		close(tasks) // Close the channel to signal no more tasks
	}()

	// Wait for all workers to complete
	go func() {
		wg.Wait()      // Wait for all worker goroutines to finish
		close(results) // Close the results channel after all workers are done
	}()

	// Collect and print results
	for result := range results {
		fmt.Println(result)
	}

	fmt.Println("All tasks have been processed.")
}
