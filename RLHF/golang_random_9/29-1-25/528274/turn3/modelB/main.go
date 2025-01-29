package main

import (
	"fmt"
	"os"
)

func processFiles(filenames []string) error {
	for _, filename := range filenames {
		file, err := os.Open(filename)
		if err != nil {
			return fmt.Errorf("error opening file %s: %w", filename, err)
		}

		// Defer the closure of the file using a function literal.
		defer func() {
			fmt.Printf("Closing file: %s\n", filename)
			if err := file.Close(); err != nil {
				fmt.Printf("Error closing file %s: %v\n", filename, err)
			}
		}()

		// Processing the file...
		// Simulate an error that occurs while processing the file
		if filename == "file2.txt" {
			return fmt.Errorf("simulated error processing file: %s", filename)
		}

		fmt.Printf("Processing file: %s\n", filename)
	}
	return nil
}

func main() {
	filenames := []string{"file1.txt", "file2.txt", "file3.txt"}
	if err := processFiles(filenames); err != nil {
		fmt.Printf("Error processing files: %v\n", err)
		return
	}
	fmt.Println("All files processed successfully.")
}
