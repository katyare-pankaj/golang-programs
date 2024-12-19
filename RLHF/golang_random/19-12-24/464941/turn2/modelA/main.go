package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(id int, work chan int, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()

	select {
	case num := <-work:
		fmt.Printf("Worker %d: processing %d\n", id, num)
		time.Sleep(time.Second)
	case <-ctx.Done():
		fmt.Printf("Worker %d: cancelled\n", id)
		return
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	numWorkers := 5
	work := make(chan int, 10) // Buffered channel
	wg := &sync.WaitGroup{}

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, work, wg, ctx)
	}

	// Send work to workers
	for i := 1; i <= 20; i++ {
		select {
		case work <- i:
		case <-ctx.Done():
			fmt.Println("Work cancelled while sending")
			return
		}
	}

	// Close the work channel to signal workers to stop
	close(work)

	// Wait for workers to finish
	wg.Wait()

	fmt.Println("All work completed")
}
