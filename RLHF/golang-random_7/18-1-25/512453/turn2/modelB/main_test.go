package main

import (
	"fmt"
	"testing"
)

// TestStringFormatting tests various scenarios of string formatting
func TestStringFormatting(t *testing.T) {
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
		{"Padding with spaces", "%10s", []interface{}{"Hello"}, "     Hello"},
		{"Padding with zeroes", "%010s", []interface{}{"Hello"}, "000000Hello"},
		{"Precision for floats", "%.3f", []interface{}{3.14159}, "3.142"},
		{"Negative integer", "%d", []interface{}{-123}, "-123"},
		{"Negative float", "%.2f", []interface{}{-3.14159}, "-3.14"},
		{"Escaped characters", "%q", []interface{}{"Hello\nWorld"}, `"Hello\nWorld"`},
		{"Invalid format specifier", "%x", []interface{}{"Hello"}, "%!x(string=Hello)"},
		// Add more test cases for different format specifiers and inputs as needed
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
