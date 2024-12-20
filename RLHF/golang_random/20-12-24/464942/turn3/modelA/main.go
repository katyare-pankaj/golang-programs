package main

import (
	"fmt"
	"os"
)

func processFiles(fileNames []string) error {
	for _, fileName := range fileNames {
		file, err := os.Open(fileName)
		if err != nil {
			return err
		}
		defer file.Close() // Ensure file is closed regardless of errors

		// Perform operations on the file
		_, err = file.WriteString("Some data\n")
		if err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}
		fmt.Println("Processed file:", fileName)
	}
	return nil
}

func main() {
	fileNames := []string{"file1.txt", "file2.txt", "file3.txt"}
	err := processFiles(fileNames)
	if err != nil {
		fmt.Println("Error processing files:", err)
	}
}
