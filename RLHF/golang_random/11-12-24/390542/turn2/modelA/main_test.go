package main

import (
	"go-programs/RLHF/golang_random/11-12-24/390542/turn1/modelA/counter"
	"sync"
	"testing"
)

func TestCounterConcurrency(t *testing.T) {
	const numGoroutines = 100
	const incrementsPerGoroutine = 1000

	c := &counter.Counter{}
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Buffer to collect results across goroutines (ensuring it's thread-safe)
	results := make([]int, numGoroutines)

	// Start multiple goroutines to increment the counter
	for i := 0; i < numGoroutines; i++ {
		go func(index int) {
			defer wg.Done()
			localCount := 0
			for j := 0; j < incrementsPerGoroutine; j++ {
				c.Increment()
				localCount++
			}
			results[index] = localCount // This write is protected by WaitGroup
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Check the final value
	expectedValue := numGoroutines * incrementsPerGoroutine
	if actualValue := c.Value(); actualValue != expectedValue {
		t.Errorf("Expected counter value %v, got %v", expectedValue, actualValue)
	}

	// Optionally, validate counts from local counts (if necessary)
	for i, count := range results {
		if count != incrementsPerGoroutine {
			t.Errorf("Goroutine %d did not increment as expected: got %d, want %d", i, count, incrementsPerGoroutine)
		}
	}
}
