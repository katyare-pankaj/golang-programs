package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulate two shared resources
var resource1 = &sync.RWMutex{}
var resource2 = &sync.RWMutex{}

// Function to access resource1
func accessResource1() {
	resource1.RLock()         // Acquire read lock on resource1
	defer resource1.RUnlock() // Release read lock on resource1

	// Simulate work
	fmt.Println("Accessing resource1...")
	time.Sleep(time.Second)
}

// Function to access resource2
func accessResource2() {
	resource2.RLock()         // Acquire read lock on resource2
	defer resource2.RUnlock() // Release read lock on resource2

	// Simulate work
	fmt.Println("Accessing resource2...")
	time.Sleep(time.Second)
}

// Function to access both resources in a consistent order
func accessResourcesConsistently() {
	resource1.RLock()         // Acquire read lock on resource1
	defer resource1.RUnlock() // Release read lock on resource1

	resource2.RLock()         // Acquire read lock on resource2
	defer resource2.RUnlock() // Release read lock on resource2

	// Simulate work
	fmt.Println("Accessing resource1 and resource2...")
	time.Sleep(time.Second)
}

func main() {
	// Create multiple goroutines to access resources
	for i := 0; i < 5; i++ {
		go accessResourcesConsistently()
		go accessResourcesConsistently()
		go accessResource1()
		go accessResource2()
	}

	// Wait for all goroutines to complete
	time.Sleep(10 * time.Second)
}
