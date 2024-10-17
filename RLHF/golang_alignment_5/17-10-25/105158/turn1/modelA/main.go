package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Define an error type
type AppError struct {
	msg string
}

func (e *AppError) Error() string {
	return e.msg
}

// Continuation passing style function to calculate area
func calculateAreaCPS(radius float64, cont func(float64, error)) {
	if radius < 0 {
		cont(0, &AppError{msg: "Radius cannot be negative"})
		return
	}

	go func() {
		area := 3.14 * radius * radius
		cont(area, nil)
	}()
}

// Using sync.Pool for object reusing
var bufferPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 1024) // Adjust size as needed
	},
}

func processDataAsync(data []byte, cont func(error)) {
	// Get a buffer from the pool
	buf := bufferPool.Get().([]byte)
	defer bufferPool.Put(buf)

	// Perform some asynchronous work (simulated with a goroutine)
	go func() {
		// Simulate work
		runtime.Gosched()
		cont(nil)
	}()
}

func main() {
	// Asynchronous usage example
	radius := 5.6
	calculateAreaCPS(radius, func(area float64, err error) {
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Area: %.2f\n", area)
	})

	data := []byte("Some data to process asynchronously")
	processDataAsync(data, func(err error) {
		if err != nil {
			fmt.Println("Processing failed:", err)
			return
		}
		fmt.Println("Data processing completed.")
	})

	// Wait for asynchronous tasks to complete
	runtime.Gosched()
	fmt.Println("Main program exiting.")
}
