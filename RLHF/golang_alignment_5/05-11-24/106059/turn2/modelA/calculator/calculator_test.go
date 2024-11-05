package calculator

import "testing"

func TestAdd(t *testing.T) {
	calc := Calculator{}
	if result := calc.Add(10.0, 20.0); result != 30.0 {
		t.Errorf("Add(10.0, 20.0) should be 30.0, got %f", result)
	}
}

func TestSubtract(t *testing.T) {
	calc := Calculator{}
	if result := calc.Subtract(30.0, 20.0); result != 10.0 {
		t.Errorf("Subtract(30.0, 20.0) should be 10.0, got %f", result)
	}
}
