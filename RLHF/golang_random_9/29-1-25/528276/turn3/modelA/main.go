package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

// writeFile is a function that writes some text to a file and ensures the file is closed properly.
func writeFile(filename string, data string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Open the file with write permission, create it if it doesn't exist.
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Failed to open file %s: %v\n", filename, err)
		return
	}

	// Ensure the file is closed when the function finishes, no matter the outcome.
	defer func() {
		if cerr := file.Close(); cerr != nil {
			log.Printf("Failed to close file %s: %v\n", filename, cerr)
		}
	}()

	// Write data to the file.
	if _, err := file.WriteString(data); err != nil {
		log.Printf("Failed to write to file %s: %v\n", filename, err)
		return
	}

	fmt.Printf("Successfully wrote to file %s\n", filename)
}

func main() {
	var wg sync.WaitGroup
	filenames := []string{
		"file1.txt",
		"file2.txt",
		"file3.txt",
	}

	for _, filename := range filenames {
		wg.Add(1)
		go writeFile(filename, "Hello, Gophers!\n", &wg)
	}

	wg.Wait()
	fmt.Println("All file operations complete.")
}
