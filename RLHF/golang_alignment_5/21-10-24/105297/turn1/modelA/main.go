package main

import (
	"fmt"
	"math/rand"
	"time"
)

// A LazyValue represents a value that is computed lazily when needed.
type LazyValue struct {
	compute  func() int
	value    int
	computed bool
}

// Get retrieves the value of the LazyValue, computing it if necessary.
func (l *LazyValue) Get() int {
	if !l.computed {
		l.value = l.compute()
		l.computed = true
	}
	return l.value
}

// newLazyValue creates a new LazyValue that will compute its value using the provided function.
func newLazyValue(compute func() int) *LazyValue {
	return &LazyValue{compute: compute}
}

// simulateApiCall represents a simulated API call that takes time and returns a random result.
func simulateApiCall() int {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	return rand.Intn(100)
}

func main() {
	// Create a lazy value that represents the result of an API call
	apiResult := newLazyValue(simulateApiCall)

	fmt.Println("Starting VR rendering...")
	// Perform various VR rendering tasks here
	time.Sleep(2 * time.Second)

	// Now, let's retrieve the value from the lazy API result
	fmt.Println("Fetching data from API...")
	data := apiResult.Get()
	fmt.Println("API Result:", data)

	// More VR rendering tasks
	time.Sleep(1 * time.Second)
	fmt.Println("VR rendering complete.")
}
