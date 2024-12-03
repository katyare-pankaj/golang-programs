package main

import (
	"fmt"
	"math"
	"math/rand"
)

// GaussianNoise adds Gaussian noise to a value with given sensitivity, epsilon, and delta
func GaussianNoise(value float64, sensitivity float64, epsilon float64, delta float64) float64 {
	// Standard deviation of the Gaussian distribution
	sigma := sensitivity * math.Sqrt(2*math.Log(1.25/delta)) / epsilon
	// Generate random noise from the Gaussian distribution
	noise := rand.NormFloat64() * sigma
	// Return the value with added noise
	return value + noise
}

func main() {
	// Initialize the random number generator
	rand.Seed(42)

	// Example data: high precision measurement
	measurement := 1234567.890123

	// Define sensitivity, epsilon, and delta
	sensitivity := 1.0 // Sensitivity depends on the context
	epsilon := 0.1
	delta := 0.01

	// Add Gaussian noise
	noisyMeasurement := GaussianNoise(measurement, sensitivity, epsilon, delta)

	// Print results
	fmt.Printf("Original Measurement: %.6f\n", measurement)
	fmt.Printf("Noisy Measurement: %.6f\n", noisyMeasurement)
}
