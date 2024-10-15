package modelB

import "testing"

func TestDivide(t *testing.T) {
	// Existing tests

	// Path 3: Division by a negative number
	_, err := divide(10, -2)
	if err == nil {
		t.Errorf("Expected an error for division by a negative number")
	}
}
