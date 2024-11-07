package main

import (
	"fmt"
)

// Vectorized multiplication using Go's built-in SIMD package.
func vectorizedMul(a, b []float64) []float64 {
	if len(a) != len(b) {
		panic("Input arrays must have the same length.")
	}

	result := make([]float64, len(a))

	// Check if SIMD is supported on the current platform.
	if simd.Supported() {
		var av, bv, rv simd.Float64x2
		for i := 0; i < len(a); i += 2 {
			av = simd.Float64x2{a[i], a[i+1]}
			bv = simd.Float64x2{b[i], b[i+1]}
			rv = av.Mul(bv)
			result[i] = rv[0]
			result[i+1] = rv[1]
		}
	} else {
		// Fall back to non-vectorized multiplication if SIMD is not supported.
		for i := range a {
			result[i] = a[i] * b[i]
		}
	}

	return result
}

func main() {

	res := vectorizedMul([]float64{10.0}, []float64{10.0})

	fmt.Println(res)
}
