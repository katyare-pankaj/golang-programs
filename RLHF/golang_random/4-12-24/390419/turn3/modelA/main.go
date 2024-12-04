package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Worker simulates a task that can take time and can respond to cancellation.
func worker(ctx context.Context, id int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	select {
	case <-time.After(time.Duration(2+id) * time.Second):
		results <- fmt.Sprintf("Worker %d completed", id)
	case <-ctx.Done():
		results <- fmt.Sprintf("Worker %d canceled", id)
	}
}

func main() {
	// Create a context with a timeout of 3 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // Ensure all resources are cleaned up

	results := make(chan string, 3)
	workerCount := 3
	var wg sync.WaitGroup

	// Start multiple worker goroutines.
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(ctx, i, results, &wg)
	}

	// Collect results or handle timeout.
	for i := 0; i < workerCount; i++ {
		select {
		case res := <-results:
			fmt.Println(res)
		case <-ctx.Done():
			fmt.Println("Main: Context timeout reached, canceling remaining workers")
			cancel() // Optionally cancel the context explicitly, though it will be canceled by the timeout
			return
		}
	}

	wg.Wait() // Wait for all workers to finish
}
