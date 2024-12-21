package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func processingRoutine(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	log.Printf("Goroutine %d started\n", id)
	for {
		select {
		case <-ctx.Done():
			log.Printf("Goroutine %d: Context canceled. Exiting!\n", id)
			return
		default:
			// Simulate some work being done
			time.Sleep(time.Second * 1)
			fmt.Printf("Goroutine %d is working...\n", id)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start a single goroutine
	wg.Add(1)
	go processingRoutine(ctx, &wg, 1)

	// Simulate some main work
	time.Sleep(time.Second * 3)
	fmt.Println("Main routine wants to cancel the goroutine now.")
	cancel()

	// Wait for the goroutine to terminate
	wg.Wait()
	fmt.Println("Program ended.")
}
