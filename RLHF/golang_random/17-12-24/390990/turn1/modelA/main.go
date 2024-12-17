package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulate a shared resource
type Resource struct {
	Name string
	mu   sync.Mutex
}

func (r *Resource) Lock() {
	fmt.Printf("Locking %s\n", r.Name)
	r.mu.Lock()
}

func (r *Resource) Unlock() {
	r.mu.Unlock()
	fmt.Printf("Unlocking %s\n", r.Name)
}

// Function that tries to acquire locks on two resources in a fixed order
func safeOperation(r1, r2 *Resource) {
	// Enforce lock ordering
	r1.Lock()
	defer r1.Unlock() // Ensure unlock happens even on panic

	// Simulate some operation
	time.Sleep(1 * time.Second)

	// After holding the first lock, try getting the second lock
	r2.Lock()
	defer r2.Unlock() // Ensure unlock happens even on panic

	// Perform operation
	fmt.Printf("Performing operation on %s and %s\n", r1.Name, r2.Name)
}

func main() {
	res1 := &Resource{Name: "Resource 1"}
	res2 := &Resource{Name: "Resource 2"}

	// Start two goroutines that will access the resources in the same order
	go safeOperation(res1, res2)
	go safeOperation(res1, res2) // Altering this second call to use res2, res1 will lead to deadlock

	// Wait for goroutines to finish
	time.Sleep(3 * time.Second)
}
