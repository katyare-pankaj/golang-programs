package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	resource1 = make(chan struct{}, 1)
	resource2 = make(chan struct{}, 1)
)

func processResource1(id int) {
	defer fmt.Printf("Goroutine %d released resource 1\n", id)

	fmt.Printf("Goroutine %d acquired resource 1\n", id)
	time.Sleep(1 * time.Second) // Simulate work
}

func processResource2(id int) {
	defer fmt.Printf("Goroutine %d released resource 2\n", id)

	fmt.Printf("Goroutine %d acquired resource 2\n", id)
	time.Sleep(1 * time.Second) // Simulate work
}

func main() {
	var wg sync.WaitGroup

	// Goroutine 1 acquires resource 1 first, then resource 2
	wg.Add(1)
	go func() {
		defer wg.Done()
		processResource1(1)
		processResource2(1)
	}()

	// Goroutine 2 acquires resource 2 first, then resource 1
	wg.Add(1)
	go func() {
		defer wg.Done()
		processResource2(2)
		processResource1(2)
	}()

	wg.Wait()
	fmt.Println("All goroutines have completed")
}
