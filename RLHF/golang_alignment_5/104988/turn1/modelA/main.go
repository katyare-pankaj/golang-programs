package main

import (
	"errors"
	"fmt"
	"testing"
)

// Example function with multiple paths
func divide(a, b int) (int, error) {
	// Added a check for negative numbers
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}

	// Check for negative divisor
	if b < 0 {
		return 0, errors.New("division by a negative number is not allowed")
	}

	result := a / b
	return result, nil
}

func main() {
	fmt.Println(divide(10, 2))  // Output: (5, <nil>)
	fmt.Println(divide(10, 0))  // Output: (0, division by zero is not allowed)
	fmt.Println(divide(-10, 2)) // Output: (0, division by a negative number is not allowed)
}

func TestDivide(t *testing.T) {
	type testCase struct {
		a      int
		b      int
		expect int
		err    string
	}
	testCases := []testCase{
		{a: 10, b: 2, expect: 5},
		{a: 10, b: 0, err: "division by zero is not allowed"},
		{a: -10, b: 2, err: "division by a negative number is not allowed"},
		{a: 0, b: -2, expect: 0},
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
