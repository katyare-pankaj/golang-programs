// Example of vectorized addition
package main

import (
	"fmt"

	"gonum.org/v1/gonum/blas"
)

func main() {
	// Create two vectors
	a := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	b := []float64{6.0, 7.0, 8.0, 9.0, 10.0}
	// Result vector
	c := make([]float64, len(a))

	// Perform vectorized addition using gonum/blas
	blas.Daxpy(len(a), 1.0, a, 1, b, 1, c, 1)

	// Print the result
	fmt.Println(c) // Output: [7.0 9.0 11.0 13.0 15.0]
}
