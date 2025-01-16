package main

import (
	"fmt"
	"os"
)

func main() {
	// Directory names
	originalDirName := "example_dir"
	newDirName := "renamed_dir"

	// Ensure the directory does not exist before creating it
	if err := os.RemoveAll(originalDirName); err != nil {
		fmt.Printf("Error removing existing directory: %v\n", err)
		return
	}

	// Create a new directory
	if err := os.Mkdir(originalDirName, 0755); err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	// Check if the directory was created successfully
	if _, err := os.Stat(originalDirName); err != nil {
		fmt.Printf("Error: Directory not created.\n")
		return
	}
	fmt.Printf("Directory '%s' created successfully.\n", originalDirName)

	// Rename the directory
	if err := os.Rename(originalDirName, newDirName); err != nil {
		fmt.Printf("Error renaming directory: %v\n", err)
		return
	}

	// Check if the directory was renamed successfully
	if _, err := os.Stat(newDirName); err != nil {
		fmt.Printf("Error: Directory not renamed.\n")
		return
	}
	fmt.Printf("Directory '%s' renamed successfully to '%s'.\n", originalDirName, newDirName)

	// Delete the directory
	if err := os.RemoveAll(newDirName); err != nil {
		fmt.Printf("Error deleting directory: %v\n", err)
		return
	}

	// Check if the directory was deleted successfully
	if _, err := os.Stat(newDirName); err != nil {
		fmt.Printf("Directory '%s' deleted successfully.\n", newDirName)
	} else {
		fmt.Printf("Error: Directory not deleted.\n")
	}
}
