package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Custom error type for more context
type fileError struct {
	err  error
	file string
}

func (e fileError) Error() string {
	return fmt.Sprintf("error processing file '%s': %v", e.file, e.err)
}

// OpenFile safely opens a file and returns a handle and an error.
func OpenFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fileError{err, filename}
	}
	return file, nil
}

// CloseFile safely closes a file handle.
func CloseFile(file *os.File) {
	if err := file.Close(); err != nil {
		fmt.Fprintf(os.Stderr, "error closing file: %v\n", err)
	}
}

// ReadFile reads a file and prints each line.
func ReadFile(filename string) error {
	file, err := OpenFile(filename)
	if err != nil {
		return err
	}
	defer CloseFile(file)

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

	if err := ReadFile(filename); err != nil {
		fmt.Fprintf(os.Stderr, "Error processing file: %v\n", err)
		os.Exit(1)
	}
}
