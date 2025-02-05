package main

import (
	"errors"
	"math/big"
	"testing"
)

// Factorial calculates the factorial of a non-negative integer n using big.Int.
func Factorial(n int64) (*big.Int, error) {
	if n < 0 {
		return nil, errors.New("negative input not allowed")
	}
	result := big.NewInt(1)
	if n == 0 || n == 1 {
		return result, nil
	}
	for i := int64(2); i <= n; i++ {
		result.Mul(result, big.NewInt(i))
	}
	return result, nil
}

func TestFactorial_PositiveInput(t *testing.T) {
	expected := big.NewInt(120)
	actual, err := Factorial(5)
	if err != nil || actual.Cmp(expected) != 0 {
		t.Errorf("Factorial(5) failed. Expected: %s, Got: %s", expected, actual)
	}
}

func TestFactorial_ZeroInput(t *testing.T) {
	expected := big.NewInt(1)
	actual, err := Factorial(0)
	if err != nil || actual.Cmp(expected) != 0 {
		t.Errorf("Factorial(0) failed. Expected: %s, Got: %s", expected, actual)
	}
}

func TestFactorial_NegativeInput(t *testing.T) {
	_, err := Factorial(-5)
	if err == nil {
		t.Errorf("Factorial(-5) should fail with an error for negative input.")
	}
}

func TestFactorial_LargeInput(t *testing.T) {
	// Test for a large number, for example, 20! = 2432902008176640000
	expected := new(big.Int)
	expected.SetString("2432902008176640000", 10)
	actual, err := Factorial(20)
	if err != nil || actual.Cmp(expected) != 0 {
		t.Errorf("Factorial(20) failed. Expected: %s, Got: %s", expected, actual)
	}
}

func main() {
	// Intentionally left empty. The main focus is on testing here.
}
