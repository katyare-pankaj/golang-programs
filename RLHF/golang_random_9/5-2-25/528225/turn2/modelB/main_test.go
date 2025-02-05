package main

import (
	"testing"
)

func Factorial(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func TestFactorial_PositiveInput(t *testing.T) {
	expected := 120
	actual := Factorial(5)
	if actual != expected {
		t.Errorf("Factorial(5) failed. Expected: %d, Got: %d", expected, actual)
	}
}

func TestFactorial_ZeroInput(t *testing.T) {
	expected := 1
	actual := Factorial(0)
	if actual != expected {
		t.Errorf("Factorial(0) failed. Expected: %d, Got: %d", expected, actual)
	}
}

func TestFactorial_NegativeInput(t *testing.T) {
	expected := 0
	actual := Factorial(-5)
	if actual != expected {
		t.Errorf("Factorial(-5) failed. Expected: %d, Got: %d", expected, actual)
	}
}

func main() {
}
