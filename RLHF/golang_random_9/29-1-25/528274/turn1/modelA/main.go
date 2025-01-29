package main

import (
	"fmt"
	"io"
	"os"
)

// processFiles demonstrates using the aggregate pattern with defer
// for efficient resource cleanup.
func processFiles(filenames []string) error {
	var files []io.Closer

	// Open files and defer their closure in case of error or completion.
	for _, filename := range filenames {
		file, err := os.Open(filename)
		if err != nil {
			return fmt.Errorf("failed to open file %s: %v", filename, err)
		}

		// Add to list of files to be closed.
		files = append(files, file)

		// Defer the closure of the file to ensure resources are cleaned up.
		defer func(f io.Closer) {
			fmt.Printf("Closing file: %s\n", filename)
			f.Close()
		}(file)

		// Assume some processing occurs here
		fmt.Printf("Processing file: %s\n", filename)
	}

	// Example of what might happen after processing
	// In reality, this would include actual file operations
	fmt.Println("Files processed successfully")

	// The deferred closures will be executed here automatically
	return nil
}

func main() {
	filenames := []string{"file1.txt", "file2.txt", "file3.txt"}

	// Process the files and handle any errors.
	if err := processFiles(filenames); err != nil {
		fmt.Println("Error processing files:", err)
		return
	}

	fmt.Println("All files processed and closed successfully")
}
