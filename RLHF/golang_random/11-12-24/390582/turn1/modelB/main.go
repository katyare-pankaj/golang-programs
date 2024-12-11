package main

import (
	"context"
	"log"
	"time"
)

func worker(ctx context.Context, id int) {
	defer log.Printf("Worker %d finished", id)

	for {
		select {
		case <-ctx.Done():
			log.Printf("Worker %d received cancellation signal: %v", id, ctx.Err())
			return
		case <-time.After(1 * time.Second):
			log.Printf("Worker %d is working...", id)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start workers
	for i := 1; i <= 3; i++ {
		go worker(ctx, i)
	}

	time.Sleep(2 * time.Second)

	log.Println("Cancelling workers...")
	cancel()

	time.Sleep(2 * time.Second) // Wait for workers to exit
}
