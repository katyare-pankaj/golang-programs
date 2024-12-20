package main

import (
	"fmt"
	"os"
)

func main() {
	// Step 1: Perform operation A
	fileA, err := os.Open("fileA.txt")
	if err != nil {
		fmt.Println("Error opening file A:", err)
		return
	}
	defer fileA.Close()
	// Perform operation A on fileA
	fmt.Println("Performed operation A on fileA")

	// Step 2: Perform operation B
	fileB, err := os.Open("fileB.txt")
	if err != nil {
		fmt.Println("Error opening file B:", err)
		return
	}
	defer fileB.Close()
	// Perform operation B on fileB
	fmt.Println("Performed operation B on fileB")
	// Step 3: Perform operation C that may fail
	if err := someOperationC(); err != nil {
		fmt.Println("Error performing operation C:", err)
		return // Avoid returning directly from between multiple defers to ensure all cleanup is handled.
	}

	fmt.Println("All operations completed successfully!")
}

func someOperationC() error {
	// Simulate error condition
	return fmt.Errorf("operation C failed")
}
