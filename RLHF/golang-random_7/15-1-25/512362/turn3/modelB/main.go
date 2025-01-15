package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Define the nested directory structure to create
	rootDir := "my_dir"
	subDir1 := "subdir1"
	subDir2 := "subdir2"
	nestedDirPath := filepath.Join(rootDir, subDir1, subDir2)

	// Create the nested directory structure
	err := os.MkdirAll(nestedDirPath, 0755)
	if err != nil {
		fmt.Printf("Error creating nested directory: %v\n", err)
		return
	}
	fmt.Println("Nested directory structure created successfully!")

	// Define the old path of the subdirectory to rename
	oldSubDirPath := filepath.Join(rootDir, subDir1)

	// Define the new path of the subdirectory after renaming
	newSubDirPath := filepath.Join(rootDir, "renamed_subdir1")

	// Rename the subdirectory
	err = os.Rename(oldSubDirPath, newSubDirPath)
	if err != nil {
		fmt.Printf("Error renaming subdirectory: %v\n", err)
		return
	}
	fmt.Println("Subdirectory renamed successfully!")

	// Remove the entire directory structure
	err = os.RemoveAll(rootDir)
	if err != nil {
		fmt.Printf("Error removing directory structure: %v\n", err)
		return
	}
	fmt.Println("Directory structure removed successfully!")
}
