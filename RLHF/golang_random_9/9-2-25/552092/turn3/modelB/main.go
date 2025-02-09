package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, channel chan<- int) {
	fmt.Printf("Worker %d is starting\n", id)
	// Do some work...
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	fmt.Printf("Worker %d has completed work\n", id)

	// Signal completion
	channel <- id
}

func main() {
	numWorkers := 5
	completedWorkers := make(chan int, numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go worker(i, completedWorkers)
	}

	// Wait for all workers to complete
	for i := 1; i <= numWorkers; i++ {
		completedWorkerId := <-completedWorkers
		fmt.Printf("Worker %d has been completed\n", completedWorkerId)
	}

	fmt.Println("All workers have completed")
}
