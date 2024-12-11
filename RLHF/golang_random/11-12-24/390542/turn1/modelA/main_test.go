package counter_test

import (
	"sync"
	"testing"

	"go-programs/RLHF/golang_random/11-12-24/390542/turn1/modelA/counter"
)

func TestCounterConcurrency(t *testing.T) {
	c := &counter.Counter{}
	var wg sync.WaitGroup

	const numGoroutines = 1000
	const incrementsPerGoroutine = 1000
	wg.Add(numGoroutines)

	// Start multiple goroutines to increment the counter
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				c.Increment()
			}
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Check the final value
	expectedValue := numGoroutines * incrementsPerGoroutine
	if c.Value() != expectedValue {
		t.Errorf("Expected counter value %v, got %v", expectedValue, c.Value())
	}
}
