package main

import (
	"fmt"
	"testing"
)

// TestSprintfNegativeAndPrecision tests fmt.Sprintf for handling negative integers
// and floating-point numbers with specified precision.
func TestSprintfNegativeAndPrecision(t *testing.T) {
	// Define a list of test cases
	testCases := []struct {
		description string
		format      string
		args        []interface{}
		expected    string
	}{
		{"Negative integer", "Negative number: %d", []interface{}{-42}, "Negative number: -42"},
		{"Negative float", "Negative float with default precision: %f", []interface{}{-3.14159}, "Negative float with default precision: -3.141590"},
		{"Negative float with precision", "Negative float: %.3f", []interface{}{-3.14159}, "Negative float: -3.142"},
		{"Negative float with more precision", "Negative float: %.5f", []interface{}{-3.14159}, "Negative float: -3.14159"},
		{"Zero float with precision", "%+.2f", []interface{}{0.0}, "+0.00"},
		{"Negative float exponential", "%.2e", []interface{}{-123456.789}, "-1.23e+05"},
		{"Mixed positive and negative values", "%d %d %.2f", []interface{}{10, -5, -2.718}, "10 -5 -2.72"},
	}

	// Iterate over the test cases
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
