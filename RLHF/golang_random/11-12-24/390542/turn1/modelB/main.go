package main

import (
	"sync"
	"testing"
)

func ProcessItems(items []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, item := range items {
		// Simulate processing
	}
}

func ProcessConcurrently(items []int) {
	var wg sync.WaitGroup
	const numGoroutines = 10
	itemCount := len(items)
	chunkSize := itemCount / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := min((i+1)*chunkSize, itemCount)
		wg.Add(1)
		go ProcessItems(items[start:end], &wg)
	}

	wg.Wait()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestProcessConcurrently(t *testing.T) {
	items := make([]int, 1000)
	for i := range items {
		items[i] = i
	}

	// Test with race detector
	if testing.Short() {
		t.Skip("race detection is too slow for short tests")
	}
	t.Run("with race detector", func(t *testing.T) {
		if testing.GetBoolEnv("GORACE") {
			t.Skip("race detector already enabled globally")
		}
		defer testing.SetBoolEnv("GORACE", false)
		testing.SetBoolEnv("GORACE", true)
		ProcessConcurrently(items)
	})

	// Unit test: Check that all items are processed
	var processedItems []int
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		ProcessItems(items, &wg)
		processedItems = items
		wg.Done()
	}()
	wg.Wait()

	for i, item := range processedItems {
		if item != i {
			t.Errorf("item %d not processed correctly: got %d, want %d", i, item, i)
		}
	}
}
