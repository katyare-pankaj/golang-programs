package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// ReadFile reads the contents of a file using defer for resource management.
func ReadFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("Error opening file: %v\n", err)
		return err
	}
	defer file.Close() // Ensure file is closed even if an error occurs

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("Error reading file: %v\n", err)
			return err
		}
		fmt.Println(string(line))
	}

	return nil
}

func main() {
	filename := "example.txt"
	if err := ReadFile(filename); err != nil {
		log.Printf("Failed to read file: %v\n", err)
	}
}
