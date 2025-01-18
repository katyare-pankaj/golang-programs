package main

import (
	"errors"
	"fmt"
	"testing"
)

// Divide takes two integers and returns the quotient and remainder.
// Returns an error if the divisor is zero.
func Divide(dividend, divisor int) (int, int, error) {
	if divisor == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder, nil
}

// TestDivide tests the Divide function for various scenarios.
func TestDivide(t *testing.T) {
	tests := []struct {
		dividend int
		divisor  int
		wantQuot int
		wantRem  int
		wantErr  bool
	}{
		{10, 2, 5, 0, false},    // 10 divided by 2 is 5 with remainder 0
		{10, 3, 3, 1, false},    // 10 divided by 3 is 3 with remainder 1
		{10, 0, 0, 0, true},     // division by zero should return an error
		{0, 5, 0, 0, false},     // 0 divided by 5 is 0 with no error
		{-10, 3, -3, -1, false}, // handling negative numbers
	}

	for _, tt := range tests {
		quot, rem, err := Divide(tt.dividend, tt.divisor)
		if (err != nil) != tt.wantErr {
			t.Errorf("Divide(%d, %d) error = %v, wantErr %v", tt.dividend, tt.divisor, err, tt.wantErr)
			continue
		}
		if quot != tt.wantQuot || rem != tt.wantRem {
			t.Errorf("Divide(%d, %d) = (%d, %d), want (%d, %d)", tt.dividend, tt.divisor, quot, rem, tt.wantQuot, tt.wantRem)
		}
	}
}

func main() {
	// Example usage of the Divide function
	quot, rem, err := Divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Quotient: %d, Remainder: %d\n", quot, rem)
	}
}
