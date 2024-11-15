package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numWorkers = 10
	maxTasks   = 1000
	taskTime   = time.Millisecond * 100
)

type task struct {
	id int
}

var wg sync.WaitGroup

func worker(ctx context.Context, workCh <-chan task, resultCh chan<- int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case task := <-workCh:
			// Simulate work
			time.Sleep(taskTime * time.Duration(rand.Intn(10)+1))
			resultCh <- task.id
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	workCh := make(chan task, numWorkers)
	resultCh := make(chan int)

	wg.Add(numWorkers)

	// Start workers
	for i := 0; i < numWorkers; i++ {
		go worker(ctx, workCh, resultCh)
	}

	// Generate and send tasks
	for i := 0; i < maxTasks; i++ {
		workCh <- task{id: i}
	}

	// Collect results
	var results []int
	for i := 0; i < maxTasks; i++ {
		select {
		case result := <-resultCh:
			results = append(results, result)
		case <-ctx.Done():
			fmt.Println("Timeout reached.")
			return
		}
	}

	fmt.Println("Results:", results)
}
