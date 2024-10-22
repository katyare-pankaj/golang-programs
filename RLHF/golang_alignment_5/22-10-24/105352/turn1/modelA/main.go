package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Initialize random number generator
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Function to perform a random calculation
func randomCalculation() int {
	a := rand.Intn(100)
	b := rand.Intn(100)
	op := rand.Intn(3) // 0 for addition, 1 for subtraction, 2 for multiplication
	var result int

	switch op {
	case 0:
		result = a + b
	case 1:
		result = a - b
	case 2:
		result = a * b
	}

	return result
}

func main() {
	// Perform 10 random calculations
	for i := 0; i < 10; i++ {
		result := randomCalculation()
		fmt.Printf("Calculation %d: %d\n", i+1, result)
	}
}
