package main

import (
	"fmt"
	"sync"
)

// Shared counter variable
var sharedCounter int = 0

// Mutex to lock access to the shared counter
var mutex sync.Mutex

func incrementCounter() {
	mutex.Lock()
	sharedCounter++
	mutex.Unlock()
}

func decrementCounter() {
	mutex.Lock()
	sharedCounter--
	mutex.Unlock()
}

func readCounter() int {
	mutex.Lock()
	defer mutex.Unlock()
	return sharedCounter
}

func main() {
	// Number of goroutines that will increment and decrement the counter
	numGoroutines := 10

	// Create a WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup

	// Start numGoroutines that will increment the counter concurrently
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			for j := 0; j < 10000; j++ {
				incrementCounter()
			}
			wg.Done()
		}()
	}

	// Start numGoroutines that will decrement the counter concurrently
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			for j := 0; j < 10000; j++ {
				decrementCounter()
			}
			wg.Done()
		}()
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Read the final value of the shared counter
	finalCounterValue := readCounter()

	fmt.Println("Final value of the shared counter:", finalCounterValue)

	// Check if the counter value is thread-safe
	if finalCounterValue != 0 {
		fmt.Println("ERROR: The shared counter is not thread-safe!")
	} else {
		fmt.Println("The shared counter is thread-safe.")
	}
}
