package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulate fetching data from an API
func fetchData(url string, duration time.Duration) []string {
	time.Sleep(duration)
	return []string{"Airport1", "Airport2", "Airport3"}
}

// Simulate rendering UI elements
func renderUI(data []string) {
	fmt.Println("Rendering UI with data:")
	for _, item := range data {
		fmt.Println(item)
	}
}

func main() {
	// We'll use a WaitGroup to wait for both asynchronous tasks to complete
	var wg sync.WaitGroup

	// Simulate API request to fetch airports
	wg.Add(1)
	go func() {
		defer wg.Done()
		airports := fetchData("https://example.com/api/airports", 2*time.Second)
		renderUI(airports)
	}()

	// Simulate other tasks running concurrently
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Running another task...")
		time.Sleep(1 * time.Second)
	}()

	// We can continue executing other code here without waiting for the tasks to finish
	fmt.Println("Main thread continues executing...")

	// Wait for both tasks to complete
	wg.Wait()
	fmt.Println("All tasks completed.")
}
