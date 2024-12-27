package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numWorkers = 5
	taskCount  = 10
)

func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d started task %d\n", id, task)
		time.Sleep(time.Second * 2) // Simulate work
		results <- task * 2
		fmt.Printf("Worker %d finished task %d\n", id, task)
	}
}

func main() {
	tasks := make(chan int, taskCount)
	results := make(chan int, taskCount)
	var wg sync.WaitGroup

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(i, tasks, results, &wg)
	}

	for i := 1; i <= taskCount; i++ {
		tasks <- i
	}

	close(tasks) // Signal workers to stop when tasks channel is closed

	// Implement timeout mechanism
	timeout := time.NewTimer(time.Second * 5) // Set a timeout of 5 seconds
	defer timeout.Stop()

	for i := 1; i <= taskCount; i++ {
		select {
		case result := <-results:
			fmt.Printf("Received result: %d\n", result)
		case <-timeout.C:
			fmt.Println("Timeout exceeded. Some tasks may not have been completed.")
			return
		}
	}

	wg.Wait() // Wait for all workers to finish
	fmt.Println("All tasks completed successfully.")
}
