package main

import (
	"fmt"
	"math/rand"
)

// LaplaceNoise adds Laplace noise to a value with a given sensitivity and epsilon
func LaplaceNoise(value float64, sensitivity float64, epsilon float64) float64 {
	// Scale parameter for the Laplace distribution
	b := sensitivity / epsilon
	// Generate random noise from the Laplace distribution
	noise := rand.ExpFloat64() * 2 * b
	// Return the value with added noise
	return value + noise
}

func main() {
	// Initialize the random number generator
	rand.Seed(42)

	// Example data: high precision measurement
	measurement := 1234567.890123

	// Define sensitivity and epsilon
	sensitivity := 1.0 // Sensitivity depends on the context
	epsilon := 0.1

	// Add Laplace noise
	noisyMeasurement := LaplaceNoise(measurement, sensitivity, epsilon)

	// Print results
	fmt.Printf("Original Measurement: %.6f\n", measurement)
	fmt.Printf("Noisy Measurement: %.6f\n", noisyMeasurement)
}
