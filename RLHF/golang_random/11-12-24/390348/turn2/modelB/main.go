package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	err := readFileAndProcess("example.txt")
	if err != nil {
		log.Fatalf("Failed to read and process file: %v", err)
	}
}

func readFileAndProcess(filename string) error {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}

	// Ensure file closure when the function exits
	defer func() {
		if cerr := file.Close(); cerr != nil {
			log.Printf("Failed to close file: %v", cerr)
		}
	}()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Process each line (in this case, simply output it)
		fmt.Println(scanner.Text())
	}

	// Check for any errors encountered while reading
	if serr := scanner.Err(); serr != nil {
		return fmt.Errorf("error reading file: %v", serr)
	}

	return nil
}
