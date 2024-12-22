package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Add the goroutines to the WaitGroup
	wg.Add(2)

	// Start the first goroutine
	go func() {
		defer wg.Done() // Decrement the counter once the goroutine finishes
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine 1 finished")
	}()

	// Start the second goroutine
	go func() {
		defer wg.Done() // Decrement the counter once the goroutine finishes
		time.Sleep(1 * time.Second)
		fmt.Println("Goroutine 2 finished")
	}()

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All goroutines finished, main function exiting")
}
