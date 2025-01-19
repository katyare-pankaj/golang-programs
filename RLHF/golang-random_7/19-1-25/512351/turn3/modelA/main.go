package main

import (
	"os"
	"path/filepath"
	"time"
)

func main() {
	// Set the target directory and max age for retention
	directory := "./data"         // Specify your directory path here
	maxAge := 30 * 24 * time.Hour // Files older than 30 days will be deleted

	// Get the current time
	now := time.Now()

	// Walk through the directory
	filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Return early if there's an error accessing a file
			return err
		}

		if !info.IsDir() {
			// Check if the file is older than the specified duration
			if now.Sub(info.ModTime()) > maxAge {
				os.Remove(path) // Delete the file
			}
		}
		return nil // Continue walking
	})
}
