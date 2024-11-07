package main

import (
	"fmt"

	"github.com/gonum/matrix/mat64"
)

func main() {
	// Sample data: A 2D matrix representing a system of linear equations
	data := [][]float64{
		{1.0, 2.0, 3.0},
		{4.0, 5.0, 6.0},
		{7.0, 8.0, 9.0},
	}
	A := mat64.NewDense(3, 3, data)

	// Sample solution vector
	b := mat64.NewVector(3, []float64{1.0, 2.0, 3.0})

	// Using gonum's LU decomposition for solving linear equations
	var lu mat64.LU
	lu.Factorize(A)
	x := mat64.NewVector(3, nil)
	lu.Solve(x, b)

	fmt.Println("Solution Vector:")
	x.Print(10, 2)
}
