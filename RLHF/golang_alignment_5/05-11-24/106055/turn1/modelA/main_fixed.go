package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

const (
	// The size of the buffer for reading/writing data
	bufferSize = 32 * 1024
)

// archiveFiles creates a gzip-compressed tar archive of the specified files.
func archiveFiles(archivePath string, filePathsChannel <-chan string, wg *sync.WaitGroup) error {
	// Open the archive file in write mode with gzip compression
	f, err := os.Create(archivePath)
	if err != nil {
		return err
	}
	defer f.Close()

	gw := gzip.NewWriter(f)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	// Process files from the channel
	for filePath := range filePathsChannel {
		if err := archiveFile(tw, filePath); err != nil {
			log.Printf("Error archiving %s: %v", filePath, err)
		}
	}

	wg.Done() // Signal that this goroutine is done
	return nil
}

func archiveFile(tw *tar.Writer, filePath string) error {
	info, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	// Create a tar header and write it to the archive
	header, err := tar.FileInfoHeader(info, "")
	if err != nil {
		return err
	}
	header.Name = filePath // Ensure the header name is set correctly

	if err := tw.WriteHeader(header); err != nil {
		return err
	}

	// Open the file for reading and write it to the archive
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := io.CopyBuffer(tw, f, make([]byte, bufferSize)); err != nil {
		return err
	}
	return nil
}

func main() {
	// Path to the archive file
	archivePath := "example.tar.gz"

	// List of files to archive
	filePaths := []string{
		"file1.txt",
		"file2.txt",
	}

	// Create a channel to send file paths to workers
	filePathsChannel := make(chan string)

	var wg sync.WaitGroup

	// Start a worker to process files
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := archiveFiles(archivePath, filePathsChannel, &wg); err != nil {
			log.Fatalf("Error creating archive: %v", err)
		}
	}()

	// Add file paths to the channel
	go func() {
		for _, filePath := range filePaths {
			filePathsChannel <- filePath
		}
		close(filePathsChannel) // Close the channel when done
	}()

	wg.Wait() // Wait for all goroutines to finish
	fmt.Printf("Archive created successfully: %s\n", archivePath)
}
