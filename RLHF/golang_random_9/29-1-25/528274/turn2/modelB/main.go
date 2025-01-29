package main

import (
	"fmt"
	"io"
	"os"
)

func processFiles(filenames []string) error {
	for _, filename := range filenames {
		// Open the file for reading and writing.
		f, err := os.OpenFile(filename, os.O_RDWR, 0666)
		if err != nil {
			return fmt.Errorf("error opening file %s: %v", filename, err)
		}

		// Defer closing the file until we have processed it.
		defer f.Close()

		// Now perform the file operations as needed. For example, you can read and write to the file using `io` functions.
		// Let's assume for simplicity that we just want to print the file contents
		content, err := io.ReadAll(f)
		if err != nil {
			return fmt.Errorf("error reading file %s: %v", filename, err)
		}
		fmt.Printf("File content: %s\n", string(content))
	}
	return nil
}
func main() {
	filenames := []string{"file1.txt", "file2.txt", "file3.txt"}

	// Pass the list of filenames to the processFiles function.
	if err := processFiles(filenames); err != nil {
		fmt.Printf("Error processing files: %v\n", err)
		return
	}
	fmt.Println("Files processed successfully.")
}
