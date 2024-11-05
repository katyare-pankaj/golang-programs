package main

import (
	"bufio"
	"compress/gzip"
	"compress/zlib"
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

func archiveData(source string, destination string, compressor string, compressionLevel int) {
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
	switch compressor {
	case "gzip":
		// Use NewWriterLevel to set the compression level for gzip
		gzWriter, err := gzip.NewWriterLevel(archiveFile, compressionLevel)
		if err != nil {
			log.Fatalf("Error creating gzip writer: %v", err)
		}
		defer gzWriter.Close()
		writer = gzWriter
	case "zlib":
		// Use NewWriterLevel to set the compression level for zlib
		zlibWriter, err := zlib.NewWriterLevel(archiveFile, compressionLevel)
		if err != nil {
			log.Fatalf("Error creating zlib writer: %v", err)
		}
		defer zlibWriter.Close()
		writer = zlibWriter
	case "lz4":
		lz4Writer := lz4.NewWriter(archiveFile)
		defer lz4Writer.Close()
		writer = lz4Writer
	default:
		log.Fatalf("Invalid compressor: %s", compressor)
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

	log.Printf("Archive created successfully: %s (Compressor: %s, Level: %d)", destination, compressor, compressionLevel)
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
	destinations := []string{
		"gzip_archive.gz",
		"lz4_archive.lz4",
	}

	// Compress with Gzip
	archiveData(source, destinations[0], "gzip", gzip.BestCompression)

	// Compress with LZ4
	archiveData(source, destinations[1], "lz4", 0) // LZ4 doesn't support levels, so use 0 or omit
}
