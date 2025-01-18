package main

import (
	"fmt"
	"testing"
)

func FormatString(s string) string {
	return fmt.Sprintf("\"%s\"", s)
}

func TestFormatString(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello", "\"hello\""},
		{"hello, world", "\"hello, world\""},
		{"", "\"\""},
		{" ", "\" \""},
		{"\\", "\"\\\""},
		{"\n", "\"\n\""},
		{"\"", "\"\\\"\""},
	}
	for _, testCase := range testCases {
		actual := FormatString(testCase.input)
		if actual != testCase.expected {
			t.Errorf("FormatString(%q) = %q, want %q", testCase.input, actual, testCase.expected)
		}
	}
}
