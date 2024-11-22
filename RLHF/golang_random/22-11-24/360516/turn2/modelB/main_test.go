package main

import (
	"sync"
	"testing"
)

var (
	// Global variable to be tested
	globalCounter int

	// Mutex to protect concurrent access to the global variable
	globalCounterMutex sync.Mutex
)

// IncrementGlobalCounterIfLessThan increments the global counter if it is less than the given limit.
func IncrementGlobalCounterIfLessThan(limit int) {
	globalCounterMutex.Lock()
	defer globalCounterMutex.Unlock()
	if globalCounter < limit {
		globalCounter++
	}
}

func TestIncrementGlobalCounterIfLessThan(t *testing.T) {
	// Test setup: Initialize the global variable to 0 before each test
	globalCounterMutex.Lock()
	globalCounter = 0
	globalCounterMutex.Unlock()

	limit := 5

	// Test 1: Increment once when counter is less than limit
	IncrementGlobalCounterIfLessThan(limit)
	globalCounterMutex.Lock()
	expected := 1
	actual := globalCounter
	globalCounterMutex.Unlock()
	if expected != actual {
		t.Errorf("Global counter not incremented correctly when counter is less than limit. Expected: %d, Actual: %d", expected, actual)
	}

	// Test 2: Increment four times when counter is less than limit
	for i := 0; i < 4; i++ {
		IncrementGlobalCounterIfLessThan(limit)
	}
	globalCounterMutex.Lock()
	expected = 5
	actual = globalCounter
	globalCounterMutex.Unlock()
	if expected != actual {
		t.Errorf("Global counter not incremented correctly multiple times. Expected: %d, Actual: %d", expected, actual)
	}

	// Test 3: Do not increment when counter is greater than or equal to limit
	IncrementGlobalCounterIfLessThan(limit)
	globalCounterMutex.Lock()
	expected = 5
	actual = globalCounter
	globalCounterMutex.Unlock()
	if expected != actual {
		t.Errorf("Global counter incremented incorrectly when counter is greater than or equal to limit. Expected: %d, Actual: %d", expected, actual)
	}
}
