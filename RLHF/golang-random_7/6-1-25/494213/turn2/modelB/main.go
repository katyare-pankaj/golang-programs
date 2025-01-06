package main

import (
	"fmt"
	"math"
)

// Error interface to handle user-defined errors
type MathError struct {
	Message string
}

func (e MathError) Error() string {
	return e.Message
}

// Sentinel value to indicate an error
const NaN = float64(math.NaN())

// Function that returns multiple values along with an error
func squareRoot(input float64) (float64, error) {
	if input < 0 {
		return NaN, MathError{"Cannot calculate square root of a negative number"}
	}
	result := math.Sqrt(input)
	return result, nil
}

func main() {
	input := -9.0
	result, err := squareRoot(input)
	if err != nil {
		//Error interface usage
		if mathErr, ok := err.(MathError); ok {
			fmt.Println("User-defined error: ", mathErr.Message)
		} else {
			fmt.Println("Unknown error: ", err.Error())
		}
	} else if result == NaN {
		//Sentinel value usage
		fmt.Println("Error: Square root of a negative number is not defined")
	} else {
		fmt.Println("Square root of ", input, " is: ", result)
	}
}
