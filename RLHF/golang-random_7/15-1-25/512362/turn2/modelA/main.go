package main

import (
	"fmt"
	"os"
)

func main() {
	// Define a path for the new directory
	dirPath := "newdir"

	// Create a directory
	err := os.Mkdir(dirPath, 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}
	fmt.Println("Directory created successfully!")

	// Define a new path for renaming the directory
	newDirPath := "updateddir"

	// Rename the directory
	err = os.Rename(dirPath, newDirPath)
	if err != nil {
		fmt.Printf("Error renaming directory: %v\n", err)
		return
	}
	fmt.Println("Directory renamed successfully!")

	// Remove the directory
	err = os.RemoveAll(newDirPath)
	if err != nil {
		fmt.Printf("Error removing directory: %v\n", err)
		return
	}
	fmt.Println("Directory removed successfully!")
}
