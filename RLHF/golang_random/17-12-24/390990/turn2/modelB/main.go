package main

import (
	"fmt"
	"sync"
)

// Example using a mutex to prevent race conditions
var counter int
var counterMutex sync.Mutex

func incrementCounterWithMutex() {
	for i := 0; i < 10000; i++ {
		counterMutex.Lock()
		counter++
		counterMutex.Unlock()
	}
}

func main() {
	// Create multiple goroutines incrementing the counter with a mutex
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			incrementCounterWithMutex()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Final counter value with mutex:", counter)

	// Example using a channel to synchronize access
	sharedData := make(chan int, 1)
	sharedData <- 0 // Initial value

	var sharedCounter int

	incrementCounterWithChannel := func() {
		for i := 0; i < 10000; i++ {
			// Wait for access
			<-sharedData

			// Increment the shared counter
			sharedCounter++

			// Release access
			sharedData <- sharedCounter
		}
	}

	// Create multiple goroutines incrementing the counter with a channel
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go incrementCounterWithChannel()
		wg.Done()
	}
	wg.Wait()

	fmt.Println("Final counter value with channel:", sharedCounter)
}
