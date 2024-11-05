package arithmetic

import (
	"errors"
)

func Add(a, b float64) (float64, error) {
	return a + b, nil
}

func Subtract(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a - b, nil
}

// Other arithmetic functions...
