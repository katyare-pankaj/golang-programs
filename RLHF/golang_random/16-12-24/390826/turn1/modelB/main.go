package main

import (
	"fmt"
	"time"
)

// Worker function that simulates some work and sends a result over a channel
func worker(id int, tasks <-chan int, results chan<- int) {
	for task := range tasks {
		fmt.Printf("Worker %d received task: %d\n", id, task)
		// Simulate work by sleeping
		time.Sleep(time.Duration(task) * time.Second)
		result := task * task
		results <- result
		fmt.Printf("Worker %d completed task: %d, result: %d\n", id, task, result)
	}
}

func main() {
	// Create channels
	tasks := make(chan int, 5)
	results := make(chan int)

	// Start worker Goroutines
	for i := 1; i <= 3; i++ {
		go worker(i, tasks, results)
	}

	// Send tasks to the tasks channel
	tasks <- 2
	tasks <- 3
	tasks <- 5
	tasks <- 4
	tasks <- 1
	close(tasks) // Close the tasks channel when all tasks are sent

	// Receive results from the results channel
	fmt.Println("Receiving results:")
	for result := range results {
		fmt.Printf("Received result: %d\n", result)
	}
}
