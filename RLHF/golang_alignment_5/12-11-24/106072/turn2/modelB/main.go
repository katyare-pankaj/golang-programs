package main

import (
	"fmt"
	"time"
)

const (
	// Maximum number of requests allowed per second
	rateLimit = 5
)

func main() {
	// Create a channel to receive requests
	requests := make(chan int)

	// Create a ticker that ticks once per second
	ticker := time.NewTicker(time.Second)

	// Go routine to process requests
	go func() {
		for range ticker.C {
			for i := 0; i < rateLimit; i++ {
				select {
				case req := <-requests:
					fmt.Println("Processing request:", req)
					// Simulate API call time
					time.Sleep(100 * time.Millisecond)
				default:
					// No more requests to process, exit the loop
					break
				}
			}
		}
	}()

	// Simulate sending requests
	for i := 0; i < 20; i++ {
		requests <- i
		// Add a small delay between requests
		time.Sleep(50 * time.Millisecond)
	}

	// Stop the ticker to exit the program
	ticker.Stop()
}
