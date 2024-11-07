package main

import (
	"fmt"
	"math"
	"runtime"

	"github.com/gonum/matrix/mat64"
	"gonum.org/v1/gonum"
)

func main() {
	// Enable AVX instructions for faster vector operations
	runtime.GOMAXPROCS(runtime.NumCPU())
	gonum.UseAVX()

	// Sample data
	data := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	vector := mat64.NewVector(len(data), data)

	// Scalar loop implementation
	scalarResult := 0.0
	for _, value := range data {
		scalarResult += math.Pow(value, 2)
	}

	// Vectorized implementation
	vectorResult := mat64.Dot(vector, vector)

	fmt.Println("Scalar Result:", scalarResult)
	fmt.Println("Vector Result:", vectorResult)
}
