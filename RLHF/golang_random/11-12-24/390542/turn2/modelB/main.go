package main

import (
	"go-programs/RLHF/golang_random/11-12-24/390542/turn2/modelB/data_processor"
	"sync"
	"testing"
)

func TestProcessDataConcurrency(t *testing.T) {
	// Enable race detector
	defer testing.Cleanup(func() {
		if r := recover(); r != nil {
			t.Errorf("panic: %v", r)
		}
	})

	// Test data
	items := []int{1, 2, 3, 4, 5}
	expectedResults := []int{1, 4, 9, 16, 25}

	// Synchronization primitives
	var wg sync.WaitGroup
	results := make(chan int, len(items))

	// Start goroutines
	wg.Add(len(items) / 2) // Example: divide work among goroutines
	for i := 0; i < len(items)/2; i++ {
		go data_processor.ProcessData(items[i*2:(i+1)*2], &wg, results)
	}
	wg.Add(len(items) % 2)
	go data_processor.ProcessData(items[len(items)/2:], &wg, results)

	// Wait for all goroutines to finish
	wg.Wait()

	// Close the results channel
	close(results)

	// Verify results
	var actualResults []int
	for result := range results {
		actualResults = append(actualResults, result)
	}

	if !sync.DeepEqual(actualResults, expectedResults) {
		t.Errorf("Expected results %v, got %v", expectedResults, actualResults)
	}
}
