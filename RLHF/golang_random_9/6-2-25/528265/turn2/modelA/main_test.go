package main

import (
	"io/ioutil"
	"os"
	"testing"
)

// Using defer for cleanup
func TestWithDefer(t *testing.T) {
	// Create a temporary file
	tmpFile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Fatal(err)
	}
	// Ensure the file is deleted when the function exits
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString("Hello, World!")
	if err != nil {
		t.Fatal(err)
	}

	// Simulate a test operation that could fail
	if err := someOperation(); err != nil {
		t.Fatal(err)
	}
}

// Manual cleanup
func TestManualCleanup(t *testing.T) {
	// Create a temporary file
	tmpFile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Fatal(err)
	}

	// Manual cleanup function
	clean := true
	cleanup := func() {
		if clean {
			os.Remove(tmpFile.Name())
		}
	}

	defer cleanup()

	_, err = tmpFile.WriteString("Hello, World!")
	if err != nil {
		clean = false
		t.Fatal(err)
	}

	// Simulate a test operation that could fail
	if err := someOperation(); err != nil {
		clean = false
		t.Fatal(err)
	}

	// Neglect to set clean to false if test completes successfully
}

func someOperation() error {
	// Simulate some operation that might fail
	return nil // Return an error here to simulate failure if needed
}

func main() {
	// To execute tests, typically you'd use: `go test`
}
