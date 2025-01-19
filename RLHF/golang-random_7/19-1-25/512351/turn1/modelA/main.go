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

// enforceDataRetention applies the retention policy by deleting files older than MaxAge
func enforceDataRetention(directory string, policy RetentionPolicy) error {
	now := time.Now()

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing file %s: %v", path, err)
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Check file age
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
	directory := "./data" // Update the directory path as needed
	policy := RetentionPolicy{
		MaxAge: 30 * 24 * time.Hour, // 30 days
	}

	err := enforceDataRetention(directory, policy)
	if err != nil {
		log.Fatalf("Error enforcing data retention: %v", err)
	}

	log.Println("Data retention enforcement complete.")
}
