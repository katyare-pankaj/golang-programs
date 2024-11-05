package main

import (
	"bufio"
	"compress/gzip"
	"io"
	"log"
	"os"
	"runtime"
	"sync"
)

const (
	chunkSize = 1024 * 1024 // 1MB chunk size
)

func archiveData(source string, destination string) {
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

	gzipWriter := gzip.NewWriter(archiveFile)
	defer gzipWriter.Close()

	// Read the source file in chunks and compress concurrently
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := compressChunks(source, gzipWriter); err != nil {
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

	reader := bufio.NewReader(file)

	for {
		chunk := make([]byte, chunkSize)
		n, err := reader.Read(chunk)
		if n == 0 {
			if err == io.EOF {
				return nil // End of file
			}
			return err // Other error
		}

		if _, err := writer.Write(chunk[:n]); err != nil {
			return err
		}
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all available CPU cores

	source := "file.txt"
	destination := "archive.gz"

	archiveData(source, destination)
}
