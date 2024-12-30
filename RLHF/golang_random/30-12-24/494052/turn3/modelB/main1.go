package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func countLines(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading lines: %v", err)
	}

	return lineCount, nil
}

func main() {
	filename := "test.txt"
	lineCount, err := countLines(filename)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	fmt.Printf("Number of lines in %s: %d\n", filename, lineCount)
}
