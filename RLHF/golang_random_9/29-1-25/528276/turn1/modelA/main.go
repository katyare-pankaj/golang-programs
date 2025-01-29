package main

import (
	"fmt"
	"net/http"
	"sync"
)

// simulateWork simulates work that a goroutine might do, including resource cleanup.
func simulateWork(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Acquiring a resource; e.g., a network connection.
	resp, err := http.Get("http://example.com/")
	if err != nil {
		fmt.Printf("Goroutine %d: Failed to get URL: %v\n", id, err)
		return
	}

	// Defer resource cleanup right after allocation.
	defer resp.Body.Close()

	fmt.Printf("Goroutine %d: Successfully got a response from example.com\n", id)
	// Additional processing...
}

func main() {
	var wg sync.WaitGroup
	const concurrency = 5

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go simulateWork(i, &wg)
	}

	wg.Wait()
	fmt.Println("All goroutines have completed!")
}
