package main

import (
	"fmt"
	"sync"
)

// Function that simulates a worker process
func worker(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range ch {
		fmt.Println("Processed:", num)
	}
}

// Main function
func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go worker(ch, &wg)

	// Sending data to the channel
	for i := 0; i < 5; i++ {
		ch <- i
	}

	// Closing the channel with error handling
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
		if ch != nil {
			closeSafely(ch)
		}
	}()

	// Close the channel
	close(ch)
	fmt.Println("Channel closed.")

	// Wait for the worker to finish processing
	wg.Wait()
}

// Function to safely close the channel and handle potential panic
func closeSafely(ch chan int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered while closing channel:", r)
		}
	}()

	// Attempt to close the channel
	close(ch) // This could panic if the channel is already closed
}
