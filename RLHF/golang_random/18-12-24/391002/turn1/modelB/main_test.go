package main

import (
	"fmt"
	"testing"
)

func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return dividend / divisor, nil
}

func worker(dividend, divisor int, result chan<- int, errChan chan<- error) {
	defer close(result)
	defer close(errChan)
	res, err := divide(dividend, divisor)
	if err != nil {
		errChan <- err
		return
	}
	result <- res
}

func TestDivide(t *testing.T) {
	tests := []struct {
		dividend, divisor int
		expectedResult    int
		expectedError     error
	}{
		{10, 2, 5, nil},
		{10, 0, 0, fmt.Errorf("division by zero")},
	}

	for _, test := range tests {
		result := make(chan int)
		errChan := make(chan error)
		go worker(test.dividend, test.divisor, result, errChan)

		select {
		case res := <-result:
			if test.expectedError != nil {
				t.Errorf("Test failed: expected error %v, got result %d", test.expectedError, res)
			} else if res != test.expectedResult {
				t.Errorf("Test failed: expected result %d, got %d", test.expectedResult, res)
			}
		case err := <-errChan:
			if test.expectedError == nil {
				t.Errorf("Test failed: expected result %d, got error %v", test.expectedResult, err)
			} else if err.Error() != test.expectedError.Error() {
				t.Errorf("Test failed: expected error %v, got %v", test.expectedError, err)
			}
		}
	}
}

func main() {
	// Running the tests
	testing.Main(func(m *testing.M) {
		return m.Run()
	}, nil, nil, nil)
}
