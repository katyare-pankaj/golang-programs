package main

import (
	"fmt"
	"time"
)

// Thunk is a function that returns a value of type int
type Thunk func() int

// CreateThunk returns a closure that defers the computation of a given function
func CreateThunk(f func() int) Thunk {
	var result *int
	return func() int {
		if result == nil {
			val := f()
			result = &val
		}
		return *result
	}
}

// ExpensiveComputation simulates a time-consuming computation
func ExpensiveComputation() int {
	fmt.Println("Performing expensive computation...")
	time.Sleep(2 * time.Second) // simulate delay
	return 42
}

func main() {
	// Creating a thunk that defers the expensive computation
	thunk := CreateThunk(ExpensiveComputation)

	fmt.Println("Thunk created, but computation not yet performed.")

	// This is where the computation is actually performed
	fmt.Println("First call to thunk:", thunk())

	// Subsequent calls to the thunk use the cached result
	fmt.Println("Second call to thunk:", thunk())
	fmt.Println("Third call to thunk:", thunk())
}
