package main

import (
	"fmt"
	"os"
)

type resource struct {
	name string
}

// openResource opens a new resource and returns its handle.
func openResource(name string) *resource {
	fmt.Printf("Opening resource: %s\n", name)
	return &resource{name}
}

// closeResource closes the specified resource.
func closeResource(r *resource) {
	fmt.Printf("Closing resource: %s\n", r.name)
}

func main() {
	resource1 := openResource("resource1")
	resource2 := openResource("resource2")

	defer closeResource(resource1)
	defer closeResource(resource2)

	// Simulate some work that might fail.
	if err := performWork(); err != nil {
		fmt.Println("Work failed:", err)
		os.Exit(1)
	}

	// If work succeeds, defer statements will be executed automatically when main() returns.
	fmt.Println("Work succeeded!")
}

func performWork() error {
	// Simulate some work that might return an error.
	return fmt.Errorf("something went wrong")
}
