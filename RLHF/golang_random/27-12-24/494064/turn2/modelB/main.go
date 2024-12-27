package main

import (
	"fmt"
	"sync"
	"time"
)

// Task defines the task that needs to be processed
type Task struct {
	ID      int
	Data    string
	Timeout time.Duration
}

func worker(id int, tasks <-chan Task, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d started task %d with data %s\n", id, task.ID, task.Data)
		select {
		case <-time.After(task.Timeout):
			fmt.Printf("Worker %d timed out for task %d\n", id, task.ID)
			results <- -1
		default:
			// Simulate work
			time.Sleep(time.Second * 2)
			fmt.Printf("Worker %d finished task %d with data %s\n", id, task.ID, task.Data)
			results <- task.ID
		}
	}
}

func main() {
	numWorkers := 5
	tasks := make(chan Task, 10)
	results := make(chan int, 10)
	var wg sync.WaitGroup

	// Create a worker pool
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(i, tasks, results, &wg)
	}

	// Add tasks to the channel with timeouts
	tasks <- Task{ID: 1, Data: "Task 1", Timeout: time.Second * 3}
	tasks <- Task{ID: 2, Data: "Task 2", Timeout: time.Second * 5}
	tasks <- Task{ID: 3, Data: "Task 3", Timeout: time.Second * 2}
	tasks <- Task{ID: 4, Data: "Task 4", Timeout: time.Second * 4}
	tasks <- Task{ID: 5, Data: "Task 5", Timeout: time.Second * 1}

	// Close the tasks channel to signal the workers to stop
	close(tasks)

	var completedTasks int
	// Wait for results and handle timeouts
	for {
		select {
		case result := <-results:
			if result == -1 {
				fmt.Println("One task timed out.")
			} else {
				completedTasks++
				fmt.Printf("Received result for task %d.\n", result)
			}
			if completedTasks == 5 {
				goto done
			}
		default:
			// Wait for all workers to finish
		}
	}

done:
	wg.Wait()
	fmt.Println("All tasks completed.")
}
