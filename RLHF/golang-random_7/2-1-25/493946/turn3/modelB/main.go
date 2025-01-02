package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d: Starting task %d\n", id, task)
		time.Sleep(time.Duration(100+task) * time.Millisecond)
		results <- task
		fmt.Printf("Worker %d: Task %d completed\n", id, task)
	}
}

func main() {
	numWorkers := 3
	numTasks := 10
	tasks := make(chan int, numTasks)
	results := make(chan int, numTasks)

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		go worker(i, tasks, results, &wg)
	}

	// Add tasks to the channel
	for i := 1; i <= numTasks; i++ {
		tasks <- i
	}
	close(tasks) // Close the task channel to signal workers that there are no more tasks

	// Wait for all workers to complete
	go func() {
		wg.Wait()
		close(results) // Close the result channel after all workers are done
	}()

	// Collect results
	for result := range results {
		fmt.Printf("Main: Task %d completed\n", result)
	}

	fmt.Println("All tasks completed. Exiting program.")
}
