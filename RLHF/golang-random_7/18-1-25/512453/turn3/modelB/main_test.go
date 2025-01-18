package main

import (
	"fmt"
	"testing"
)

func TestSprintfFormatting(t *testing.T) {
	testCases := []struct {
		description string
		format      string
		args        []interface{}
		expected    string
	}{
		{"Negative integer", "%d", []interface{}{-123}, "-123"},
		{"Negative integer with precision", "%.2d", []interface{}{-123}, "-12"},
		{"Positive integer", "%d", []interface{}{123}, "123"},
		{"Positive integer with precision", "%.2d", []interface{}{123}, "12"},
		{"Floating-point number", "%.2f", []interface{}{3.14159}, "3.14"},
		{"Negative floating-point number", "%.2f", []interface{}{-3.14159}, "-3.14"},
		{"Floating-point number with zero precision", "%.0f", []interface{}{3.14159}, "3"},
		{"Negative floating-point number with zero precision", "%.0f", []interface{}{-3.14159}, "-3"},
		{"Floating-point number with large precision", "%.10f", []interface{}{3.14159}, "3.1415900000"},
		{"Negative floating-point number with large precision", "%.10f", []interface{}{-3.14159}, "-3.1415900000"},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := fmt.Sprintf(tc.format, tc.args...)
			if actual != tc.expected {
				t.Errorf("Expected %q, but got %q", tc.expected, actual)
			}
		})
	}
}
