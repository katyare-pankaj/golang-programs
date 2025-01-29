package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func processFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Simulate some work with the file
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}
	fmt.Println("File content:", string(data))

	return nil
}

func main() {
	err := processFile("example.txt")
	if err != nil {
		fmt.Println("Error processing file:", err)
	} else {
		fmt.Println("File processing completed successfully.")
	}
}
