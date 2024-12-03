package main

import (
	"fmt"
	"sync"
	"time"
)

// Resources represents shared resources protected by mutexes
type Resources struct {
	mu1 sync.Mutex
	mu2 sync.Mutex
}

// AccessResource1 safely accesses resource 1
func (r *Resources) AccessResource1() {
	r.mu1.Lock() // Acquire lock for resource 1
	defer r.mu1.Unlock()

	fmt.Println("Accessing resource 1")
	time.Sleep(100 * time.Millisecond)

	// Access resource 2
	r.AccessResource2()
}

// AccessResource2 safely accesses resource 2
func (r *Resources) AccessResource2() {
	r.mu2.Lock() // Acquire lock for resource 2
	defer r.mu2.Unlock()

	fmt.Println("Accessing resource 2")
	time.Sleep(100 * time.Millisecond)
}

func main() {
	resources := Resources{}

	go resources.AccessResource1()
	go resources.AccessResource1()

	time.Sleep(200 * time.Millisecond) // Wait for the resources to be accessed

	fmt.Println("All operations completed")
}
