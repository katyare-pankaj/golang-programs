package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: Exiting due to timeout\n", id)
			return
		default:
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Printf("Worker %d: Working...\n", id)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 3
	wg.Add(numWorkers)
	// Create a context with a timeout of 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	for i := 0; i < numWorkers; i++ {
		go worker(ctx, &wg, i)
	}
	wg.Wait()
	fmt.Println("All workers have finished or timed out.")
}
