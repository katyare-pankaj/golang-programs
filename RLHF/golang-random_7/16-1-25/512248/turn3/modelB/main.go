package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type FileProcessor struct {
	Filename string
}

// NewFileProcessor creates a new FileProcessor
func NewFileProcessor(filename string) *FileProcessor {
	return &FileProcessor{Filename: filename}
}

// ProcessFile reads lines from the file and outputs them to the console.
func (fp *FileProcessor) ProcessFile() error {
	file, err := os.Open(fp.Filename)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		fmt.Println(line)
	}
	return scanner.Err()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <filename>\n", os.Args[0])
		os.Exit(1)
	}

	filename := os.Args[1]
	fp := NewFileProcessor(filename)
	if err := fp.ProcessFile(); err != nil {
		handleError(err)
	}
}

// handleError prints the error message to stderr and exits with a non-zero status.
func handleError(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(1)
}
