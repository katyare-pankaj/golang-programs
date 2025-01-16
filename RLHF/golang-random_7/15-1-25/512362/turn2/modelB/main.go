package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	baseDir := "example_dir"

	// Create a new directory using filepath.Join to build the path
	newDirPath := filepath.Join(baseDir, "newdir")
	err := os.MkdirAll(newDirPath, 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}
	fmt.Println("Directory created successfully!")

	// Rename the directory using filepath.Base to extract the base name
	oldDirPath := newDirPath
	newDirBase := filepath.Base(newDirPath) + "_updated"
	newDirPath = filepath.Join(filepath.Dir(oldDirPath), newDirBase)
	err = os.Rename(oldDirPath, newDirPath)
	if err != nil {
		fmt.Printf("Error renaming directory: %v\n", err)
		return
	}
	fmt.Println("Directory renamed successfully!")

	// Remove the directory using filepath.Join to build the path (ignoring walkErr if not nil)
	defer os.RemoveAll(newDirPath)
	err = filepath.Walk(newDirPath, func(path string, info os.FileInfo, walkErr error) error {
		return nil
	})
	if err != nil {
		fmt.Printf("Error removing directory: %v\n", err)
	} else {
		fmt.Println("Directory removed successfully!")
	}
}
