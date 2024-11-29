package main

import (
	"context"
	"fmt"
	"time"
)

func work(ctx context.Context, id int) {
	defer fmt.Printf("Worker %d completed\n", id)
	select {
	case <-ctx.Done():
		fmt.Printf("Worker %d: Context canceled\n", id)
		return
	case <-time.After(10 * time.Second):
		fmt.Printf("Worker %d: Doing work...\n", id)
	}
}

func main() {
	numWorkers := 5
	workerCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for i := 0; i < numWorkers; i++ {
		go work(workerCtx, i)
	}

	select {
	case <-workerCtx.Done():
		fmt.Println("Work completed within the timeout period.")
	default:
		fmt.Println("Work exceeded the timeout period.")
	}
}
