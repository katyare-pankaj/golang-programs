package main

import "testing"

func TestDivide_EdgeCases(t *testing.T) {
	testCases := []struct {
		a      int
		b      int
		expect int
		err    string
	}{
		{a: 10, b: 0, err: "division by zero is not allowed"}, // Existing test
		{a: 0, b: 2, expect: 0},                               // New test for division by zero numerator
		{a: 10, b: -2, err: "division by a negative number is not allowed"},
		{a: -10, b: -2, expect: 5},
		{a: IntMax, b: 2, expect: IntMax / 2},  // Test for integer overflow
		{a: IntMin, b: -2, expect: IntMin / 2}, // Test for integer underflow
	}

	for _, tc := range testCases {
		result, err := divide(tc.a, tc.b)
		if tc.err != "" {
			if err == nil || err.Error() != tc.err {
				t.Errorf("Expected error '%s', got '%v' for input (%d, %d)", tc.err, err, tc.a, tc.b)
			}
		} else if result != tc.expect {
			t.Errorf("Expected %d, got %d for input (%d, %d)", tc.expect, result, tc.a, tc.b)
		}
	}
}
