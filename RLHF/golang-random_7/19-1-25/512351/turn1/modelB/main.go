package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type retentionPolicy struct {
	Age     time.Duration
	Pattern string
}

var dataRetentionPolicies = []retentionPolicy{
	{
		Age:     7 * 24 * time.Hour,
		Pattern: "*log*.txt",
	},
	{
		Age:     30 * 24 * time.Hour,
		Pattern: "*backup*.tar.gz",
	},
}

func applyRetentionPolicies(dataDir string) error {
	now := time.Now()
	for _, policy := range dataRetentionPolicies {
		if err := filepath.Walk(dataDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return fmt.Errorf("error walking %q: %v", path, err)
			}
			if info.IsDir() {
				return nil
			}
			match, err := filepath.Match(policy.Pattern, filepath.Base(path))
			if err != nil {
				return fmt.Errorf("error matching pattern %q for %q: %v", policy.Pattern, path, err)
			}
			if !match {
				return nil
			}
			if now.Sub(info.ModTime()) >= policy.Age {
				fmt.Printf("Deleting file: %s (modified %s, age %s)\n", path, info.ModTime(), now.Sub(info.ModTime()))
				if err := os.Remove(path); err != nil {
					return fmt.Errorf("error removing %q: %v", path, err)
				}
			}
			return nil
		}); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	dataDir := "./data" // Replace this with the actual path to your data directory
	err := applyRetentionPolicies(dataDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error applying retention policies: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Retention policies applied successfully.")
}
