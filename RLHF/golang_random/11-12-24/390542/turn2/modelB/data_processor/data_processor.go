package data_processor

import (
	"sync"
	"time"
)

// ProcessData processes data items concurrently.
func ProcessData(items []int, wg *sync.WaitGroup, results chan int) {
	defer wg.Done()
	for _, item := range items {
		// Simulate some work
		time.Sleep(time.Millisecond * 1)
		results <- item * item // Example: square the item
	}
}
