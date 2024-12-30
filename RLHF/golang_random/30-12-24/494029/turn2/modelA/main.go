package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	// Number of goroutines to create
	numGoroutines := 10

	// Function for each goroutine to execute
	printMessage := func(message string, id int) {
		defer wg.Done()
		time.Sleep(time.Duration(id) * time.Millisecond)
		fmt.Printf("Goroutine %d: %s\n", id, message)
	}

	// Start goroutines
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go printMessage(fmt.Sprintf("Hello, World! %d", i+1), i+1)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	fmt.Println("All goroutines have finished.")
}
