package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func processData(ctx context.Context, data string) error {
	defer wg.Done()
	_, ok := ctx.Deadline()
	if !ok {
		log.Printf("No deadline set for data: %s", data)
		return nil
	}

	duration := time.Duration(rand.Intn(10)) * time.Second

	select {
	case <-ctx.Done():
		log.Printf("Data processing for %s canceled: %v", data, ctx.Err())
		return ctx.Err()
	case <-time.After(duration):
		log.Printf("Data processing for %s complete in %d seconds", data, duration.Seconds())
		return nil
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Create a basic context
	ctx := context.Background()

	// Create a context with a deadline
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	// Function to process some data
	go func() {
		if err := processData(ctx, "Data1"); err != nil {
			fmt.Println("Error:", err)
		}
	}()

	// Function to process some data
	go func() {
		if err := processData(ctx, "Data2"); err != nil {
			fmt.Println("Error:", err)
		}
	}()

	wg.Add(2) // Two goroutines for processing

	// Wait for all goroutines to complete
	wg.Wait()

	fmt.Println("Main function finished.")
}
