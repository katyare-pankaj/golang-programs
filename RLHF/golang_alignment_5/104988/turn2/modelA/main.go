package main

import (
	"fmt"
	"testing"
)

// Example function with multiple paths
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}

	// Add a check for negative divisor
	if b < 0 {
		return 0, fmt.Errorf("division by a negative number is not allowed")
	}

	return a / b, nil
}

func TestDivide_EdgeCases(t *testing.T) {
	type testCase struct {
		a      int
		b      int
		expect int
		err    string
	}
	testCases := []testCase{
		// Existing test cases
		{a: 0, b: 1, expect: 0},
		{a: 1, b: 1, expect: 1},
		{a: -1, b: 1, expect: -1},
		{a: IntMax, b: 1, expect: IntMax},
		{a: IntMin, b: 1, expect: IntMin},
		{a: IntMax, b: -1, expect: IntMin}, // Overflow
		{a: IntMin, b: -1, expect: IntMax}, // Overflow
		{a: 0, b: -1, err: "division by a negative number is not allowed"},
	}

	for _, tc := range testCases {
		result, err := divide(tc.a, tc.b)
		if tc.err != "" {
			if err == nil || err.Error() != tc.err {
				t.Errorf("Expected error '%s', got '%v'", tc.err, err)
			}
		} else if result != tc.expect {
			t.Errorf("Expected %d, got %d", tc.expect, result)
		}
	}
}

const (
	IntMax = int(^uint(0) >> 1)
	IntMin = -IntMax - 1
)

func main() {
	fmt.Println(divide(12, 2)) // Output: (5, <nil>)
	fmt.Println(divide(5, 0))  // Output: (0, division by zero is not allowed)
	fmt.Println(divide(65, 2)) // Output: (0, division by a negative number is not allowed)
}
