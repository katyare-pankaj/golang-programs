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

// IncrementConditionally increments the global counter by 1 if it is less than 10.
func IncrementConditionally() {
	globalCounterMutex.Lock()
	defer globalCounterMutex.Unlock()

	if globalCounter < 10 {
		globalCounter++
	}
}

func TestIncrementConditionally(t *testing.T) {
	// Test setup: Initialize the global variable to 0 before each test
	globalCounterMutex.Lock()
	globalCounter = 0
	globalCounterMutex.Unlock()

	// Test case 1: Counter is initially 0
	IncrementConditionally()
	globalCounterMutex.Lock()
	expected := 1
	actual := globalCounter
	globalCounterMutex.Unlock()
	if expected != actual {
		t.Errorf("Global counter not incremented correctly. Expected: %d, Actual: %d", expected, actual)
	}

	// Test case 2: Counter is initially 9
	globalCounterMutex.Lock()
	globalCounter = 9
	globalCounterMutex.Unlock()
	IncrementConditionally()
	globalCounterMutex.Lock()
	expected = 10
	actual = globalCounter
	globalCounterMutex.Unlock()
	if expected != actual {
		t.Errorf("Global counter not incremented correctly. Expected: %d, Actual: %d", expected, actual)
	}

	// Test case 3: Counter is initially 10
	globalCounterMutex.Lock()
	globalCounter = 10
	globalCounterMutex.Unlock()
	IncrementConditionally()
	globalCounterMutex.Lock()
	expected = 10
	actual = globalCounter
	globalCounterMutex.Unlock()
	if expected != actual {
		t.Errorf("Global counter not incremented correctly. Expected: %d, Actual: %d", expected, actual)
	}

	// Test case 4: Counter is initially negative
	globalCounterMutex.Lock()
	globalCounter = -5
	globalCounterMutex.Unlock()
	IncrementConditionally()
	globalCounterMutex.Lock()
	expected = -4
	actual = globalCounter
	globalCounterMutex.Unlock()
	if expected != actual {
		t.Errorf("Global counter not incremented correctly. Expected: %d, Actual: %d", expected, actual)
	}
}
