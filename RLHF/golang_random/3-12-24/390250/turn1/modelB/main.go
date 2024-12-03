package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Laplace mechanism implementation
func laplaceMechanism(trueCount int, sensitivity float64, epsilon float64) float64 {
	// Calculate the scale of the Laplace noise
	scale := sensitivity / epsilon

	// Generate Laplace noise
	rand.Seed(time.Now().UnixNano())
	noise := rand.Laplace(0, scale)

	// Add noise to the true count
	noisyCount := float64(trueCount) + noise
	return noisyCount
}

// Laplace random number generator
func randLaplace(mean float64, scale float64) float64 {
	u := rand.Float64() - 0.5
	return mean - scale*math.Signbit(u)*math.Log(1-2*math.Abs(u))
}

func main() {
	// Example usage
	trueCount := 100   // The true count of items
	sensitivity := 1.0 // Sensitivity for count queries is typically 1
	epsilon := 0.5     // Privacy parameter

	noisyCount := laplaceMechanism(trueCount, sensitivity, epsilon)
	fmt.Printf("True Count: %d, Noisy Count: %.2f\n", trueCount, noisyCount)
}
