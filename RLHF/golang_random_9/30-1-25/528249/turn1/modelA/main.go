package main

import (
	"fmt"
	"log"
	"os"
)

// Example function that opens a file and ensures it is closed
func processFile(filename string) error {
	// Open the file
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	// Ensures the file is closed upon function exit
	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("error closing file: %v", err)
		}
	}()

	// Simulate processing of the file
	fmt.Println("Processing file:", filename)

	return nil // No error
}

// Example function that acquires a resource and ensures it is released
func doSomeWork() error {
	// Simulating acquiring a resource (e.g., lock, network connection, etc.)
	fmt.Println("Acquiring resource")
	defer fmt.Println("Releasing resource") // Deferred resource release

	// Simulate doing some work
	fmt.Println("Doing some work")

	return nil // No error
}

func main() {
	// Call processFile function and handle errors
	if err := processFile("example.txt"); err != nil {
		log.Fatalf("Error processing file: %v", err)
	}

	// Call doSomeWork function and handle errors
	if err := doSomeWork(); err != nil {
		log.Fatalf("Error in doing work: %v", err)
	}

	fmt.Println("Program completed successfully")
}
