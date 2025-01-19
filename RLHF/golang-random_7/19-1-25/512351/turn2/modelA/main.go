package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// RetentionPolicy defines the rules for file retention
type RetentionPolicy struct {
	MaxAge time.Duration // Maximum age a file can have before it should be deleted
}

// enforceDataRetention deletes files older than MaxAge in the given directory
func enforceDataRetention(directory string, policy RetentionPolicy) error {
	now := time.Now()

	// Walk through the directory
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing file %s: %v", path, err)
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Calculate file age
		fileAge := now.Sub(info.ModTime())
		if fileAge > policy.MaxAge {
			log.Printf("Deleting file: %s, Age: %v\n", path, fileAge)
			err := os.Remove(path)
			if err != nil {
				return fmt.Errorf("unable to delete file %s: %v", path, err)
			}
		}
		return nil
	})

	return err
}

func main() {
	directory := "./data" // Update this path to the target directory
	policy := RetentionPolicy{
		MaxAge: 30 * 24 * time.Hour, // 30 days retention policy
	}

	// Enforce data retention policy
	err := enforceDataRetention(directory, policy)
	if err != nil {
		log.Fatalf("Error enforcing data retention: %v", err)
	}

	log.Println("Data retention enforcement complete.")
}
