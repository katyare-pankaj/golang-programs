package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, id int, ch chan<- error) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: Stopped\n", id)
			return
		default:
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			if rand.Intn(10) == 0 {
				err := errors.New("Worker failed")
				fmt.Printf("Worker %d: Failed with error: %v\n", id, err)
				ch <- err
				return
			}
			fmt.Printf("Worker %d: Working...\n", id)
		}
	}
}
func main() {
	numWorkers := 5
	wg := &sync.WaitGroup{}
	ch := make(chan error, numWorkers)
	ctx, cancel := context.WithCancel(context.Background())
	// Add a timeout to the context
	ctx, _ = context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(ctx, wg, i, ch)
	}
	go func() {
		for err := range ch {
			fmt.Printf("Received error: %v\n", err)
			cancel() // Cancel all worker Goroutines
			break
		}
	}()
	wg.Wait()
	// Check if the context was canceled due to timeout
	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("Timed out waiting for workers to finish.")
	} else {
		fmt.Println("All workers finished.")
	}
}
