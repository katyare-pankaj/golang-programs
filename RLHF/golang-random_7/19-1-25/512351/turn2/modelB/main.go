package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const (
	// The directory to monitor for retention policy enforcement.
	dataDir = "./data"
	// The duration after which files should be deleted.
	retentionDuration = 30 * 24 * time.Hour
	// The number of goroutines to use for parallel file processing.
	concurrency = 10
	// The log file path for deletion errors.
	errorLogFile = "deletion_errors.log"
)

// enforceDataRetention iterates through files in the directory and deletes those older than the specified duration.
func enforceDataRetention(dir string, duration time.Duration, errorLog *log.Logger) {
	var wg sync.WaitGroup
	wg.Add(concurrency)

	fileChan := make(chan string, concurrency)

	go func() {
		wg.Wait()
		close(fileChan)
	}()

	// Start worker goroutines
	for i := 0; i < concurrency; i++ {
		go worker(fileChan, duration, errorLog)
	}

	// Recursively walk through the directory and send file paths to the worker goroutines
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing file %s: %v", path, err)
		}

		if !info.IsDir() && time.Since(info.ModTime()) > duration {
			wg.Done()
			fileChan <- path
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Error while walking directory: %v", err)
	}
}

func worker(fileChan chan string, duration time.Duration, errorLog *log.Logger) {
	for filePath := range fileChan {
		if err := os.Remove(filePath); err != nil {
			errorLog.Printf("Error deleting file %s: %v", filePath, err)
		} else {
			log.Printf("Deleted file: %s (older than %v)\n", filePath, duration)
		}
	}
}

func main() {
	// Create a logger for deletion errors
	errorLogFile, err := os.OpenFile(errorLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening error log file: %v", err)
	}
	defer errorLogFile.Close()
	errorLog := log.New(errorLogFile, "", log.LstdFlags)

	// Check if the data directory exists
	if _, err := os.Stat(dataDir); err != nil {
		log.Fatalf("Error accessing data directory: %v", err)
	}

	enforceDataRetention(dataDir, retentionDuration, errorLog)

	log.Println("Data retention enforcement complete.")
}
