package main

import (
	"fmt"
	"io"
	"os"
)

func processFiles(filenames []string) error {
	var files []io.Closer // Slice to store opened files for closing

	// Ensure all files are closed on function exit
	defer func() {
		for _, file := range files {
			if err := file.Close(); err != nil {
				fmt.Printf("Error closing file: %v\n", err)
			} else {
				fmt.Println("File closed successfully")
			}
		}
	}()

	for _, filename := range filenames {
		file, err := os.Open(filename)
		if err != nil {
			return fmt.Errorf("failed to open file %s: %v", filename, err)
		}
		fmt.Printf("Opened file: %s\n", filename)

		// Append the file handle to the slice
		files = append(files, file)

		// Simulate some file processing
		content, err := io.ReadAll(file)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %v", filename, err)
		}

		fmt.Printf("Contents of %s: %s\n", filename, string(content))
	}

	return nil
}

func main() {
	filenames := []string{"file1.txt", "file2.txt", "file3.txt"}

	if err := processFiles(filenames); err != nil {
		fmt.Printf("Error processing files: %v\n", err)
	} else {
		fmt.Println("All files processed successfully.")
	}
}
