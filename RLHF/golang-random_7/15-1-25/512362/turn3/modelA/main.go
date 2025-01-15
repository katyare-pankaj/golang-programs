package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Define the root path for the nested directory structure
	rootPath := "nested_structure"

	// Define the structure of the directories
	dirStructure := []string{"level1", "level2", "level3"}

	// Function to create a nested directory structure
	createNestedDirs := func(path string, dirs []string) error {
		for _, dir := range dirs {
			newPath := filepath.Join(path, dir)
			if err := os.Mkdir(newPath, 0755); err != nil {
				return fmt.Errorf("error creating directory %s: %v", newPath, err)
			}
		}
		return nil
	}

	// Create the nested directory structure
	if err := createNestedDirs(rootPath, dirStructure); err != nil {
		fmt.Printf("Error creating nested directories: %v\n", err)
		return
	}
	fmt.Println("Nested directory structure created successfully!")

	// Define the subdirectory to rename
	subDirToRename := "level2"
	newSubDirName := "updated_level2"

	// Function to rename a directory
	renameDir := func(oldPath, newPath string) error {
		return os.Rename(oldPath, newPath)
	}

	// Rename the subdirectory
	oldSubDirPath := filepath.Join(rootPath, subDirToRename)
	newSubDirPath := filepath.Join(rootPath, newSubDirName)
	if err := renameDir(oldSubDirPath, newSubDirPath); err != nil {
		fmt.Printf("Error renaming directory: %v\n", err)
		return
	}
	fmt.Println("Subdirectory renamed successfully!")

	// Function to delete the entire directory structure
	removeDirStructure := func(path string) error {
		return os.RemoveAll(path)
	}

	// Remove the entire directory structure
	if err := removeDirStructure(rootPath); err != nil {
		fmt.Printf("Error removing directory structure: %v\n", err)
		return
	}
	fmt.Println("Directory structure removed successfully!")
}
