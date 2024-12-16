package main

import (
	"fmt"
	"sync"
)

func increment(counterChannel chan<- int) {
	for i := 0; i < 1000; i++ {
		counterChannel <- 1 // Send an increment signal
	}
}

func main() {
	var counter = 0
	var wg sync.WaitGroup
	counterChannel := make(chan int, 1000) // Buffered channel

	// Start 10 goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(counterChannel)
	}

	// Goroutine to process increments
	go func() {
		for increment := range counterChannel {
			counter += increment
		}
		close(counterChannel)
		wg.Done()
	}()

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("Final counter value:", counter)
}
