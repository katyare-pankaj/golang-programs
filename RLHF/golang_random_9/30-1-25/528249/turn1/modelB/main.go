package main

import (
	"fmt"
	"io"
	"os"
)

// WriteFile writes data to a file and returns an error if any.
func WriteFile(filename string, data []byte) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	return err
}

// ReadFile reads data from a file and returns it as a byte slice.
func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func main() {
	filename := "example.txt"
	dataToWrite := []byte("Hello, world!")

	// Write data to the file using WriteFile function
	err := WriteFile(filename, dataToWrite)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Read data from the file using ReadFile function
	dataRead, err := ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	// Display the read data
	fmt.Println("Data read from file:", string(dataRead))
}
