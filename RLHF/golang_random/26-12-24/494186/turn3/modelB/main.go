package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker", id, "received termination signal")
			return
		default:
			fmt.Println("Worker", id, "is working...")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	numWorkers := 3

	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go worker(ctx, i, &wg)
	}

	// Graceful shutdown handling
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Received termination signal. Gracefully shutting down...")
		cancel()
		wg.Wait()
		fmt.Println("All goroutines stopped. Exiting.")
		os.Exit(0)
	}()

	// Main application logic here
	fmt.Println("Main program is running...")
	time.Sleep(5 * time.Second)
}
