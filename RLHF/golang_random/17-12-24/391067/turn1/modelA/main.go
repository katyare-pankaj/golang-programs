package main

import (
	"fmt"
	"log"
	"time"
)

// Function to handle panic in a goroutine
func safeExecute(f func()) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
		}
	}()

	// Call the function that may panic
	f()
}

// A function that will cause a panic
func riskyFunction(id int) {
	if id%2 == 0 { // Simulate a condition for panic
		panic(fmt.Sprintf("Panic triggered in goroutine %d", id))
	}
	fmt.Printf("Goroutine %d completed successfully\n", id)
}

func main() {
	// Create a slice of goroutines
	for i := 0; i < 10; i++ {
		go safeExecute(func() func() {
			i := i // capture current value of i
			return func() { riskyFunction(i) }
		}())
	}

	// Allow time for goroutines to finish
	time.Sleep(1 * time.Second)
	fmt.Println("Main function finished.")
}
