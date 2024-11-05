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
func archiveFiles(archivePath string, filePaths []string) error {
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
	var wg sync.WaitGroup

	// Create a worker pool to process files concurrently
	const numWorkers = 4
	for i := 0; i < numWorkers; i++ {
		go func() {
			for {
				select {
				case filePath := <-filePathsChannel:
					if err := archiveFile(tw, filePath); err != nil {
						log.Printf("Error archiving %s: %v", filePath, err)
					}
				default:
					return
				}
			}
		}()
	}

	// Close the channel to signal workers to finish
	close(filePathsChannel)
	wg.Wait()

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
		"directory/file3.txt",
		"directory/subdirectory/file4.txt",
		// Add more file paths as needed
	}

	// Create a channel to send file paths to workers
	filePathsChannel := make(chan string)

	// Add file paths to the channel
	go func() {
		for _, filePath := range filePaths {
			filePathsChannel <- filePath
		}
	}()

	if err := archiveFiles(archivePath, filePathsChannel); err != nil {
		log.Fatalf("Error creating archive: %v", err)
	}

	fmt.Printf("Archive created successfully: %s\n", archivePath)
}
