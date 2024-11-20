package main

import (
	"fmt"
	"time"
)

func scheduledFunction(done chan bool) {
	for {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("Scheduled task running...")
		case <-done:
			fmt.Println("Scheduled task stopped.")
			return
		}
	}
}

func scheduleWithErrorHandling() error {
	done := make(chan bool)
	go scheduledFunction(done)
	time.Sleep(5 * time.Second) // Simulate some work being done
	close(done)

	return nil // You can return errors if anything unexpected happens
}

func main() {
	err := scheduleWithErrorHandling()
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
	}
	fmt.Println("Program completed.")
}
