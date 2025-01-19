package main

import (
	"fmt"
	"time"
)

// Callback defines the type for the callback function
type Callback func(int)

// highFrequencyOperation executes the callback in a tight loop
func highFrequencyOperation(callback Callback, iterations int) {
	for i := 0; i < iterations; i++ {
		callback(i)
	}
}

func main() {
	start := time.Now()

	// Define a simple callback function
	callback := func(n int) {
		// Simulate some work
		_ = n * n
	}

	// Run the high-frequency operation with the callback
	highFrequencyOperation(callback, 10000000) // 10 million iterations

	fmt.Printf("Operation took %s\n", time.Since(start))
}
