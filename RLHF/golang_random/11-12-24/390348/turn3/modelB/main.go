package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

func main() {
	files := []string{"file1.txt", "file2.txt", "file3.txt"}
	var wg sync.WaitGroup

	for _, filename := range files {
		wg.Add(1)
		go func(fname string) {
			defer wg.Done()
			processFile(fname)
		}(filename)
	}

	wg.Wait()
}

func processFile(filename string) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("Failed to open file %s: %v", filename, err)
		return
	}

	// Ensure file closure
	defer func() {
		if cerr := file.Close(); cerr != nil {
			log.Printf("Failed to close file %s: %v", filename, cerr)
		}
	}()

	fmt.Printf("Processing file: %s\n", filename)
	// Simulate file processing
	// This could involve reading, modifying, or analyzing the file content
}
