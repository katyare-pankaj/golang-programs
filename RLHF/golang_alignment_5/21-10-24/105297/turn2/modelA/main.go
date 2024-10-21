package main

import (
	"fmt"
	"math/rand"
	"time"
)

// LazyMap represents a lazy map that applies a function to each element only when needed.
type LazyMap struct {
	source []int
	fn     func(int) int
}

// NewLazyMap creates a new LazyMap.
func NewLazyMap(source []int, fn func(int) int) *LazyMap {
	return &LazyMap{source: source, fn: fn}
}

// Get applies the function to each element lazily and returns the resulting slice.
func (lm *LazyMap) Get() []int {
	result := make([]int, len(lm.source))
	for i, v := range lm.source {
		result[i] = lm.fn(v)
	}
	return result
}

func main() {
	// Generate a large slice of random numbers
	const size = 1_000_000
	rand.Seed(time.Now().UnixNano())
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Intn(1000)
	}

	// Define a heavy computation function
	heavyComputation := func(x int) int {
		time.Sleep(time.Microsecond * 100) // Simulate computation time
		return x * x
	}

	// Create a lazy map without applying the computation
	lazyResult := NewLazyMap(data, heavyComputation)

	fmt.Println("Lazy Map Created. Memory Usage: Low") // Memory usage is low at this point since computation is lazy

	// Now, let's perform some operations on the lazy map
	sum := 0
	for _, v := range lazyResult.Get() {
		sum += v
	}

	fmt.Println("Sum:", sum)
	fmt.Println("Computation Complete. Memory Usage: High (after applying computation)")
}
