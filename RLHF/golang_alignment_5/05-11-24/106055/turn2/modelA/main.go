package main

import (
	"compress/gzip"
	"io"
	"log"
	"os"
	"runtime"
	"sync"

	"github.com/pierrec/lz4"
)

const (
	chunkSize = 1024 * 1024 // 1MB chunk size
)

func archiveData(source string, destination string, compressionType string) {
	// Check if source file exists
	_, err := os.Stat(source)
	if os.IsNotExist(err) {
		log.Fatalf("Error: Source file '%s' not found.", source)
	}

	// Create the archive file
	archiveFile, err := os.Create(destination)
	if err != nil {
		log.Fatalf("Error creating archive file: %v", err)
	}
	defer archiveFile.Close()

	var writer io.Writer
	switch compressionType {
	case "gzip":
		writer = gzip.NewWriter(archiveFile)
		defer writer.(*gzip.Writer).Close()
	case "lz4":
		writer = lz4.NewWriter(archiveFile)
		defer writer.(*lz4.Writer).Close()
	default:
		log.Fatalf("Invalid compression type: %s", compressionType)
		return
	}

	// Read the source file in chunks and compress concurrently
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := compressChunks(source, writer); err != nil {
			log.Fatalf("Error compressing data: %v", err)
		}
	}()

	wg.Wait()

	log.Printf("Archive created successfully: %s", destination)
}

func compressChunks(source string, writer io.Writer) error {
	file, err := os.Open(source)
	if err != nil {
		return err
	}
	defer file.Close()

	buffer := make([]byte, chunkSize)

	for {
		n, err := file.Read(buffer)
		if n == 0 {
			if err == io.EOF {
				return nil // End of file
			}
			return err // Other error
		}

		if _, err := writer.Write(buffer[:n]); err != nil {
			return err
		}
	}
}

func main() {
	source := "file.txt"
	destinations := []string{
		"gzip_archive.gz",
		"lz4_archive.lz4",
	}
	compressionTypes := []string{"gzip", "lz4"}

	for i, compressionType := range compressionTypes {
		runtime.GOMAXPROCS(runtime.NumCPU()) // Use all available CPU cores

		destination := destinations[i]

		archiveData(source, destination, compressionType)
	}
}
