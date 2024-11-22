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

// IncrementGlobalCounter increments the global counter by 1.
func IncrementGlobalCounter() {
	globalCounterMutex.Lock()
	defer globalCounterMutex.Unlock()
	globalCounter++
}

func TestIncrementGlobalCounter(t *testing.T) {
	// Test setup: Initialize the global variable to 0 before each test
	globalCounterMutex.Lock()
	globalCounter = 0
	globalCounterMutex.Unlock()

	// Call the function that increments the global counter
	IncrementGlobalCounter()

	// Verify that the global counter has been incremented by 1
	globalCounterMutex.Lock()
	expected := 1
	actual := globalCounter
	globalCounterMutex.Unlock()

	if expected != actual {
		t.Errorf("Global counter not incremented correctly. Expected: %d, Actual: %d", expected, actual)
	}
}
