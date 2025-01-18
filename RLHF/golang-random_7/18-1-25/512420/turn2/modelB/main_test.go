package main

import (
	"fmt"
	"testing"
)

// SafeDivide calculates the quotient and remainder of two integers.
// It returns an error if the divisor is zero.
func SafeDivide(dividend, divisor int) (int, int, error) {
	if divisor == 0 {
		return 0, 0, fmt.Errorf("cannot divide by zero")
	}
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder, nil
}

// TestSafeDivide tests the SafeDivide function.
func TestSafeDivide(t *testing.T) {
	testCases := []struct {
		dividend, divisor int
		wantQuot, wantRem int
		wantErr           bool
	}{
		{10, 2, 5, 0, false},
		{10, 3, 3, 1, false},
		{10, 0, 0, 0, true},
		{0, 5, 0, 0, false},
		{-10, 3, -3, -1, false},
	}
	for _, tc := range testCases {
		quot, rem, err := SafeDivide(tc.dividend, tc.divisor)
		if tc.wantErr {
			if err == nil {
				t.Errorf("SafeDivide(%d, %d): expected error, got nil", tc.dividend, tc.divisor)
			}
		} else {
			if err != nil {
				t.Errorf("SafeDivide(%d, %d): unexpected error: %v", tc.dividend, tc.divisor, err)
			} else if quot != tc.wantQuot || rem != tc.wantRem {
				t.Errorf("SafeDivide(%d, %d): got (%d, %d), want (%d, %d)", tc.dividend, tc.divisor, quot, rem, tc.wantQuot, tc.wantRem)
			}
		}
	}
}

func main() {
	fmt.Println("Running tests...")
	TestSafeDivide(&testing.T{})
	fmt.Println("Tests completed successfully!")
}
