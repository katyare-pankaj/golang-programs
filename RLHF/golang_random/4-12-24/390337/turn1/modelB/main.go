package main

import (
	"fmt"
	"log"
	"runtime/debug"
	"time"
)

func main() {
	// Starting multiple concurrent tasks
	for i := 0; i < 5; i++ {
		go safeGo(func() {
			// This could panic for some index
			if i == 2 {
				panic(fmt.Sprintf("A critical error occurred in goroutine %d", i))
			}
			// Simulate some work
			time.Sleep(2 * time.Second)
			fmt.Printf("Goroutine %d completed successfully\n", i)
		})
	}

	// Waiting for some arbitrary time for demonstration purpose
	time.Sleep(5 * time.Second)
	fmt.Println("Main program completed.")
}

// safeGo executes a function in a new goroutine and recovers from panic.
func safeGo(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered from panic: %v\n", r)
				log.Printf("Stack trace: \n%s", string(debug.Stack()))
				// Optionally, notify or re-throw the panic to interrupt program if required.
			}
		}()
		fn()
	}()
}
