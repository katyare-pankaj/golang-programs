package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int, work chan int, result chan int) {
	for n := range work {
		fmt.Printf("Worker %d received task: %d\n", id, n)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		result <- n * n
		fmt.Printf("Worker %d completed task: %d\n", id, n)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	numWorkers := 3
	numTasks := 10

	// Create channels
	work := make(chan int)
	result := make(chan int)
	wg := &sync.WaitGroup{}

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, work, result)
	}

	// Send tasks
	for i := 1; i <= numTasks; i++ {
		work <- i
	}

	// Close work channel to signal workers to stop
	close(work)

	// Collect results
	for i := 1; i <= numTasks; i++ {
		res := <-result
		fmt.Printf("Result: %d\n", res)
	}

	wg.Wait()
	close(result)
}
