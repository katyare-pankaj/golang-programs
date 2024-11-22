package main

import (
	"testing"
)

// Global variable
var counter int

func TestCounter(t *testing.T) {
	// Reset the global variable before each test
	resetCounter()

	t.Run("Increment", func(t *testing.T) {
		incrementCounter()
		if counter != 1 {
			t.Errorf("expected counter to be 1, got %d", counter)
		}
	})

	t.Run("Decrement", func(t *testing.T) {
		incrementCounter()
		decrementCounter()
		if counter != 0 {
			t.Errorf("expected counter to be 0, got %d", counter)
		}
	})
}

// Helper function to increment the global counter
func incrementCounter() {
	counter++
}

// Helper function to decrement the global counter
func decrementCounter() {
	counter--
}

// Helper function to reset the global counter
func resetCounter() {
	counter = 0
}
