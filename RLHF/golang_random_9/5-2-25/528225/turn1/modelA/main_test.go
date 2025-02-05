package main

import (
	"testing"
)

// Test for Factorial function
func TestFactorial(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"Factorial of 0", 0, 1},
		{"Factorial of 1", 1, 1},
		{"Factorial of 5", 5, 120},
		{"Factorial of 3", 3, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Factorial(tt.n)
			if got != tt.want {
				t.Errorf("Factorial(%d) = %d; want %d", tt.n, got, tt.want)
			}
		})
	}
}
