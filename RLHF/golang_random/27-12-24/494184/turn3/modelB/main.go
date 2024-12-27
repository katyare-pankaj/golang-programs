package main

import (
	"fmt"
	"sync"
	"time"
)

// Task defines the task that needs to be processed
type Task struct {
	ID   int
	Data string
}

func worker(id int, tasks <-chan Task, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d started task %d with data %s\n", id, task.ID, task.Data)
		time.Sleep(time.Second * 2) // Simulate work
		fmt.Printf("Worker %d finished task %d with data %s\n", id, task.ID, task.Data)
		results <- task.ID
	}
}

func main() {
	numWorkers := 5
	tasks := make(chan Task, 10)
	results := make(chan int, 10)
	var wg sync.WaitGroup

	// Create a worker pool with a fixed number of workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	// Add tasks to the channel
	tasks <- Task{ID: 1, Data: "Task 1"}
	tasks <- Task{ID: 2, Data: "Task 2"}
	tasks <- Task{ID: 3, Data: "Task 3"}
	tasks <- Task{ID: 4, Data: "Task 4"}
	tasks <- Task{ID: 5, Data: "Task 5"}
	tasks <- Task{ID: 6, Data: "Task 6"}
	tasks <- Task{ID: 7, Data: "Task 7"}
	tasks <- Task{ID: 8, Data: "Task 8"}
	tasks <- Task{ID: 9, Data: "Task 9"}
	tasks <- Task{ID: 10, Data: "Task 10"}

	// Close the tasks channel to signal the workers to stop
	close(tasks)

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All tasks completed.")

	// Process results
	close(results)
	for result := range results {
		fmt.Printf("Received result for task %d.\n", result)
	}
}
