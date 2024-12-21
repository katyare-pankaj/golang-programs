package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Task struct represents a task to be processed by a worker
type Task struct {
	ID       int
	Name     string
	Duration time.Duration // time taken to process the task
}

// Worker simulates a worker processing a task
func Worker(id int, taskChannel <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range taskChannel {
		time.Sleep(task.Duration)
		fmt.Printf("Worker %d processed task %d: %s\n", id, task.ID, task.Name)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Number of tasks
	numTasks := 20

	// Create a buffered channel to hold tasks
	taskChannel := make(chan Task, 100) // Adjust buffer size as needed

	// Number of workers
	numWorkers := 5

	// Create a WaitGroup
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go Worker(i, taskChannel, &wg)
	}

	// Send tasks to the channel
	for i := 1; i <= numTasks; i++ {
		task := Task{
			ID:       i,
			Name:     fmt.Sprintf("Task %d", i),
			Duration: time.Duration(rand.Intn(5)+1) * time.Second,
		}
		taskChannel <- task
	}

	// Close the channel to signal the workers that no more tasks will be sent
	close(taskChannel)

	// Wait for all worker goroutines to complete
	wg.Wait()

	fmt.Println("\nAll tasks have been processed by the workers.")
}
