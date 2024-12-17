package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func divideByZero() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Recovered from panic in divideByZero: %v\n", err)
			// Optionally, print the stack trace
			buf := make([]byte, 1024*64)
			runtime.Stack(buf, false)
			log.Printf("Stack trace:\n%s\n", string(buf))
		}
	}()
	// Trigger a panic
	panic("division by zero")
}

func main() {
	fmt.Println("Starting main program...")

	// Create a goroutine that will panic
	go divideByZero()

	// Simulate some work in the main goroutine
	time.Sleep(2 * time.Second)

	fmt.Println("Main program continuing...")

	// The main program will continue running even if the goroutine panics
}
