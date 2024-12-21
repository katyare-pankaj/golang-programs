package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	select {
	case <-ctx.Done():
		fmt.Printf("Worker %d canceled\n", id)
		return
	default:
		// Simulate work
		time.Sleep(time.Duration(id) * time.Second)
		fmt.Printf("Worker %d completed\n", id)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(ctx, &wg, i)
	}

	// Simulate a timeout scenario
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout reached, canceling workers...")
		cancel()
	case <-wg.Done():
		fmt.Println("All workers completed")
	}

	// Wait for all workers to finish
	wg.Wait()
}
