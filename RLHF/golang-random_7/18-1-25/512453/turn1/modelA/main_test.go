package main

import (
	"fmt"
	"testing"
)

// TestFmtSprintf tests various scenarios of fmt.Sprintf
func TestFmtSprintf(t *testing.T) {
	// Define test cases
	testCases := []struct {
		description string
		format      string
		args        []interface{}
		expected    string
	}{
		{"Basic string", "Hello, %s", []interface{}{"World"}, "Hello, World"},
		{"Integer", "Number: %d", []interface{}{123}, "Number: 123"},
		{"Float", "Pi is approximately %.2f", []interface{}{3.14159}, "Pi is approximately 3.14"},
		{"Multiple arguments", "%s %d %0.1f", []interface{}{"Test", 42, 3.5}, "Test 42 3.5"},
		{"Edge case: empty string", "", []interface{}{}, ""},
		{"Edge case: large number", "%d", []interface{}{1234567890123456789}, "1234567890123456789"},
		// Add more edge cases as needed
	}

	// Iterate through the test cases
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// Use fmt.Sprintf to format the string
			actual := fmt.Sprintf(tc.format, tc.args...)
			// Compare the actual output with the expected output
			if actual != tc.expected {
				t.Errorf("Expected %q, but got %q", tc.expected, actual)
			}
		})
	}
}
