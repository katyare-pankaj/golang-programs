package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

// Simulates a long-running operation
func longRunningOperation(ctx context.Context, filePath string) error {
	defer log.Println("Cleaning up file handle")

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	select {
	case <-time.After(5 * time.Second): // Simulate 5 seconds of work
		fmt.Println("Operation completed successfully")
		return nil
	case <-ctx.Done():
		log.Println("Operation cancelled")
		return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	filePath := "example.txt"

	go func() {
		if err := longRunningOperation(ctx, filePath); err != nil {
			log.Printf("Error in longRunningOperation: %v\n", err)
		}
	}()

	time.Sleep(2 * time.Second)
	log.Println("Cancelling operation...")
	cancel() // Cancel the context

	select {
	case <-time.After(1 * time.Second):
		log.Println("Main thread exiting")
	}
}
