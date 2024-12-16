package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int) // Create a channel for communication

	wg.Add(10) // We will launch 10 goroutines

	// Launching 10 goroutines
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// Send the result to the channel
			ch <- 1 // Each goroutine sends '1' to the channel
		}()
	}

	go func() {
		wg.Wait() // Wait for all goroutines to finish
		close(ch) // Close the channel after all sends are done
	}()

	sum := 0
	// Receiving values from the channel
	for count := range ch {
		sum += count // Accumulate the received values
	}

	fmt.Println("Final counter value:", sum) // Should consistently output 10
}
