package main

import (
	"fmt"
	"testing"
)

func divide(dividend int, divisor int) (int, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	if dividend == 0 {
		return 0, fmt.Errorf("dividend cannot be zero")
	}
	return dividend / divisor, nil
}

func TestDivide(t *testing.T) {
	testCases := []struct {
		dividend int
		divisor  int
		wantQuot int
		wantErr  error
	}{
		{10, 2, 5, nil},
		{0, 2, 0, fmt.Errorf("dividend cannot be zero")},
		{10, 0, 0, fmt.Errorf("division by zero")},
		{-10, 2, -5, nil},
	}

	for _, tt := range testCases {
		t.Run(fmt.Sprintf("%d/%d", tt.dividend, tt.divisor), func(t *testing.T) {
			gotQuot, gotErr := divide(tt.dividend, tt.divisor)
			if gotQuot != tt.wantQuot {
				t.Errorf("divide() quotient = %v, want %v", gotQuot, tt.wantQuot)
			}
			if gotErr == nil && tt.wantErr != nil || gotErr != nil && tt.wantErr == nil || gotErr.Error() != tt.wantErr.Error() {
				t.Errorf("divide() error = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
