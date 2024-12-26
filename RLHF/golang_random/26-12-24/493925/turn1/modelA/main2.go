package main

import (
	"fmt"
)

func worker(id int, tasks <-chan int, done chan<- bool) {
	for task := range tasks {
		fmt.Println("Worker", id, "processing task", task)
		// Simulate work
		for i := 0; i < 10; i++ {
			fmt.Println("Worker", id, "doing work...")
		}
	}
	fmt.Println("Worker", id, "finished processing.")
	done <- true
}

func main() {
	tasks := make(chan int, 5)
	done := make(chan bool, 3)

	for i := 1; i <= 3; i++ {
		go worker(i, tasks, done)
	}

	// Send tasks
	fmt.Println("Main sending tasks...")
	tasks <- 1
	tasks <- 2
	tasks <- 3
	tasks <- 4
	tasks <- 5
	close(tasks)

	// Wait for workers to complete
	fmt.Println("Main starting to wait for workers...")
	for i := 1; i <= 3; i++ {
		<-done
	}
	fmt.Println("Main finished waiting.")
}
