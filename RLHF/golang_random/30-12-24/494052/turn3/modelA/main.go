package main

import (
	"fmt"
	"log"
	"os"
)

func countLines(filename string) int {
	var lineCount int
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file %s: %v", filename, err)
	}
	defer file.Close()

	scanner := os.NewScanner(file)
	for scanner.Scan() {
		lineCount++
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning file %s: %v", filename, err)
	}

	return lineCount
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: go run count_lines.go <filename>")
	}
	filename := os.Args[1]

	lineCount := countLines(filename)
	fmt.Printf("Number of lines in %s: %d\n", filename, lineCount)
}
