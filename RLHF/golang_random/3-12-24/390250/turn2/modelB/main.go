package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// LaplaceNoise adds Laplace noise to a value with a given sensitivity and epsilon
func LaplaceNoise(value float64, sensitivity float64, epsilon float64) float64 {
	// Calculate the scale of the Laplace distribution
	b := sensitivity / epsilon

	// Generate two uniformly random numbers
	u1 := rand.Float64()
	u2 := rand.Float64()

	// Use inverse transform sampling to get Laplace noise
	sign := 1.0
	if u2 < 0.5 {
		sign = -1.0
	}
	noise := sign * math.Log(1-u1) * b

	// Add the noise to the original value
	return value + noise
}

func main() {
	// Seed random number generator for reproducibility
	rand.Seed(time.Now().UnixNano())

	// Original data: a numerical value
	originalValue := 12345.6789

	// Define sensitivity and epsilon
	sensitivity := 1.0 // Sensitivity of the query
	epsilon := 0.5     // Privacy parameter

	// Add Laplace noise to the original value
	noisyValue := LaplaceNoise(originalValue, sensitivity, epsilon)

	// Print the original and noisy values
	fmt.Printf("Original Value: %.4f\n", originalValue)
	fmt.Printf("Noisy Value: %.4f\n", noisyValue)
}
