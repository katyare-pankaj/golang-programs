package main

import (
	"fmt"
	"log"
	"sync"
)

// Safe execute a function in a goroutine with recovery
func executeWithRecovery(wg *sync.WaitGroup, f func()) {
	defer wg.Done() // Notify the WaitGroup when done
	defer func() {  // Recover from panic
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
		}
	}()

	// Run the provided function
	f()
}

// Risky function that could panic
func riskyFunction(id int) {
	if id%3 == 0 { // Simulate a condition for panic every third goroutine
		panic(fmt.Sprintf("panic in goroutine %d", id))
	}
	fmt.Printf("Goroutine %d completed successfully\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1) // Increment WaitGroup counter
		go executeWithRecovery(&wg, func(i int) func() {
			return func() {
				riskyFunction(i)
			}
		}(i))
	}

	// Wait for all goroutines to complete
	wg.Wait()
	fmt.Println("All goroutines finished.")
}
