package main

import (
	"fmt"
	"testing"
)

// TestSprintf tests the fmt.Sprintf function with various input scenarios.
func TestSprintf(t *testing.T) {
	// Define a list of test cases
	testCases := []struct {
		description string
		format      string
		args        []interface{}
		expected    string
	}{
		{"Format simple string", "Hello, %s", []interface{}{"World"}, "Hello, World"},
		{"Format integer", "The answer is %d", []interface{}{42}, "The answer is 42"},
		{"Format float", "Pi is %.2f", []interface{}{3.14159}, "Pi is 3.14"},
		{"Multiple placeholders", "%s scored %d out of %d", []interface{}{"Alice", 90, 100}, "Alice scored 90 out of 100"},
		{"Edge case: empty string", "", []interface{}{}, ""},
		{"Edge case: zero values", "%d %s %.1f", []interface{}{0, "", 0.0}, "0  0.0"},
		{"Edge case: special characters", "Special: %% %v", []interface{}{nil}, "Special: % <nil>"},
		// Add more cases as needed
	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// Use fmt.Sprintf to format the string
			actual := fmt.Sprintf(tc.format, tc.args...)

			// Check if the actual output matches the expected output
			if actual != tc.expected {
				t.Errorf("For %s: expected %q, got %q", tc.description, tc.expected, actual)
			}
		})
	}
}
