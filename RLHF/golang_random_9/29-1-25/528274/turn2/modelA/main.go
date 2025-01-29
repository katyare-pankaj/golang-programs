package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func processMultipleFiles(filenames []string) error {
	// We'll use a slice to store all open files, which will be closed later.
	var files []io.Closer

	// Ensure that all files are properly closed on function exit.
	defer func() {
		for _, file := range files {
			fmt.Printf("Closing file\n")
			file.Close()
		}
	}()

	for _, filename := range filenames {
		file, err := os.Open(filename)
		if err != nil {
			return fmt.Errorf("failed to open file %s: %w", filename, err)
		}

		fmt.Printf("Opened file: %s\n", filename)

		// Append the file to the list of files
		files = append(files, file)

		// Perform file processing (example: read the first line)
		if err := readFirstLine(file); err != nil {
			return fmt.Errorf("failed to process file %s: %w", filename, err)
		}
	}

	fmt.Println("All files processed successfully")
	return nil
}

func readFirstLine(file io.Reader) error {
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		fmt.Println("Read line:", scanner.Text()) // Process the line (example)
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading line: %w", err)
	}
	return nil
}

func main() {
	filenames := []string{"file1.txt", "file2.txt", "file3.txt"}

	if err := processMultipleFiles(filenames); err != nil {
		fmt.Println("Error processing files:", err)
	} else {
		fmt.Println("All operations completed successfully")
	}
}
