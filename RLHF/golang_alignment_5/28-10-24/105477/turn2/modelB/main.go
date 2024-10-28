package main

import (
	"fmt"
	"io"
	"os"
)

// Use pre-allocated error object
var errEOF = io.EOF

func readAndPrint(r io.Reader) error {
	buf := make([]byte, 4096)
	for {
		n, err := r.Read(buf)
		if err != nil {
			// Handle error efficiently
			if err == errEOF {
				return nil // End of file, no error
			}
			return fmt.Errorf("read error: %w", err)
		}
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	f, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	// Inline error handling
	if err := readAndPrint(f); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
