package main

import (
	"fmt"
)

// Define an interface for callbacks
type Callback interface {
	Call()
}

// Register a new callback function
func RegisterCallback(cb Callback) {
	callbacks = append(callbacks, cb)
}

// Execute all registered callbacks
func ExecuteCallbacks() {
	for _, cb := range callbacks {
		cb.Call()
	}
}

var callbacks []Callback

func main() {
	// Registering a simple callback function
	RegisterCallback(func() {
		fmt.Println("Hello from a dynamic callback!")
	})

	// Registering another callback function
	RegisterCallback(func() {
		fmt.Println("Hello from another dynamic callback!")
	})

	// Execute all registered callbacks
	ExecuteCallbacks()
}
