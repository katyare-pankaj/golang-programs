package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

// Example of using defer for cleanup
func TestDeferCleanup(t *testing.T) {
	// Create a temporary file
	tempFile, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}

	// Defer the deletion of the temporary file
	defer os.Remove(tempFile.Name())

	// Write some data to the temporary file
	data := []byte("Hello, world!")
	_, err = tempFile.Write(data)
	if err != nil {
		t.Fatalf("Error writing to temporary file: %v", err)
	}

	// Read the data from the temporary file
	readData, err := ioutil.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Error reading from temporary file: %v", err)
	}

	// Compare the read data with the original data
	if string(readData) != string(data) {
		t.Errorf("Data read from temporary file does not match original data")
	}
}

// Example of manual cleanup
func TestManualCleanup(t *testing.T) {
	// Create a temporary file
	tempFile, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}

	// Manual cleanup of the temporary file
	defer func() {
		if err := os.Remove(tempFile.Name()); err != nil {
			t.Errorf("Error deleting temporary file: %v", err)
		}
	}()

	// Write some data to the temporary file
	data := []byte("Hello, world!")
	_, err = tempFile.Write(data)
	if err != nil {
		t.Fatalf("Error writing to temporary file: %v", err)
	}

	// Read the data from the temporary file
	readData, err := ioutil.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Error reading from temporary file: %v", err)
	}

	// Compare the read data with the original data
	if string(readData) != string(data) {
		t.Errorf("Data read from temporary file does not match original data")
	}
}

func main() {
	fmt.Println("Running tests...")
	// Run the test functions
	testing.Main()
	fmt.Println("Tests completed.")
}
