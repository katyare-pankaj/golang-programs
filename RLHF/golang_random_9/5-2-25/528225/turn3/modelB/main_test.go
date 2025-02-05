package main

import (
	"math"
	"testing"
)

func Factorial(n float64) float64 {
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return 1
	}
	result := 1.0
	for i := 2; i <= int(n); i++ {
		result *= float64(i)
	}
	return result
}

func TestFactorial_PositiveIntegerInput(t *testing.T) {
	expected := 120.0
	actual := Factorial(5.0)
	if actual != expected {
		t.Errorf("Factorial(5) failed. Expected: %v, Got: %v", expected, actual)
	}
}

func TestFactorial_ZeroInput(t *testing.T) {
	expected := 1.0
	actual := Factorial(0.0)
	if actual != expected {
		t.Errorf("Factorial(0) failed. Expected: %v, Got: %v", expected, actual)
	}
}

func TestFactorial_NegativeInput(t *testing.T) {
	expected := 0.0
	actual := Factorial(-5.0)
	if actual != expected {
		t.Errorf("Factorial(-5) failed. Expected: %v, Got: %v", expected, actual)
	}
}

func TestFactorial_LargeIntegerInput(t *testing.T) {
	// Use math.Inf() to handle large input
	expected := math.Inf(1)
	actual := Factorial(math.Inf(1))
	if actual != expected {
		t.Errorf("Factorial(LargeNumber) failed. Expected: %v, Got: %v", expected, actual)
	}
}

func TestFactorial_FloatingPointInput(t *testing.T) {
	// Test for fractional input, it should return 0
	expected := 0.0
	actual := Factorial(3.14)
	if actual != expected {
		t.Errorf("Factorial(FloatNumber) failed. Expected: %v, Got: %v", expected, actual)
	}
}

func TestFactorial_MinimumIntegerValue(t *testing.T) {
	// Test for the minimum integer value that the function can handle
	expected := 0.0
	actual := Factorial(-math.MaxInt64)
	if actual != expected {
		t.Errorf("Factorial(-9223372036854775808) failed. Expected: %v, Got: %v", expected, actual)
	}
}

func main() {
}
