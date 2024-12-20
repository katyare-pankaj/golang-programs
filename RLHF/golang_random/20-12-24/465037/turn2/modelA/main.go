package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func concurrentWorker(ctx context.Context, wg *sync.WaitGroup, workID int, workDuration time.Duration) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		fmt.Printf("Worker %d: Canceled or timed out\n", workID)
		return
	default:
		fmt.Printf("Worker %d: Started\n", workID)
		time.Sleep(workDuration)
		fmt.Printf("Worker %d: Completed\n", workID)
	}
}

func main() {
	numWorkers := 5
	wg := &sync.WaitGroup{}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	for i := 0; i < numWorkers; i++ {
		workDuration := time.Duration(rand.Intn(5)) * time.Second
		wg.Add(1)
		go concurrentWorker(ctx, wg, i, workDuration)
	}

	wg.Wait()
	fmt.Println("All workers completed or timed out")
}
