package main

import "fmt"

// Define a callback interface with a single method that takes an integer and returns an integer.
type Callback interface {
	Apply(int) int
}

// Implement the callback interface by defining a function that matches the interface signature.
func MyCallback(x int) int {
	return x * x
}

// Function that takes a callback as an argument and calls it with the argument value.
func ApplyCallback(callback Callback, value int) int {
	return callback.Apply(value)
}

func main() {
	// Create a variable of type Callback and assign it the function MyCallback.
	var cb Callback = MyCallback

	// Call the ApplyCallback function with the callback variable and the value 5.
	result := ApplyCallback(cb, 5)

	fmt.Println("Result:", result) // Output: Result: 25
}
