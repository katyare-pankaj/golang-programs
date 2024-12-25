package main

import (
	"fmt"
	"sync"
	"time"
)

// Shared resource
type SharedResource struct {
	data []int
}

// Initialize the shared resource
func initializeResource() *SharedResource {
	fmt.Println("Initializing shared resource...")
	time.Sleep(1 * time.Second) // Simulate initialization time
	return &SharedResource{data: []int{1, 2, 3, 4, 5}}
}

func main() {
	var once sync.Once
	var sharedResource *SharedResource

	// Function to get the shared resource
	getSharedResource := func() *SharedResource {
		once.Do(func() {
			sharedResource = initializeResource()
		})
		return sharedResource
	}

	// Multiple goroutines trying to access the shared resource
	numGoroutines := 10
	for i := 0; i < numGoroutines; i++ {
		go func() {
			resource := getSharedResource()
			fmt.Printf("Goroutine %d accessed resource: %v\n", i, resource.data)
		}()
	}

	// Wait for all goroutines to finish
	time.Sleep(2 * time.Second)
}
