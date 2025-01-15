package main

import (
	"fmt"
	"os"
)

func main() {
	// Create a directory called "newdir"
	newDirPath := "newdir"
	err := os.Mkdir(newDirPath, 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}
	fmt.Println("Directory created successfully!")

	// Rename the directory "newdir" to "updateddir"
	oldDirPath := newDirPath
	newDirPath = "updateddir"
	err = os.Rename(oldDirPath, newDirPath)
	if err != nil {
		fmt.Printf("Error renaming directory: %v\n", err)
		return
	}
	fmt.Println("Directory renamed successfully!")

	// Remove the directory "updateddir"
	err = os.RemoveAll(newDirPath)
	if err != nil {
		fmt.Printf("Error removing directory: %v\n", err)
		return
	}
	fmt.Println("Directory removed successfully!")
}
