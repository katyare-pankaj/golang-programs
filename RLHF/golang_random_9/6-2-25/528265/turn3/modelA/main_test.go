package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

// Using defer to clean up resources
func TestDeferCleanup(t *testing.T) {
	// Create a temporary file
	tmpFile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Fatal(err)
	}
	// Defer the cleanup action
	defer os.Remove(tmpFile.Name())

	// Perform some operations that might cause early returns
	_, err = tmpFile.WriteString("Testing defer.")
	if err != nil {
		t.Fatal(err) // Cleanup with defer still happens
	}

	// Further operations...
	fmt.Println("Continuing test with defer cleanup.")
}

// Manual resource cleanup example
func TestManualCleanup(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Fatal(err)
	}

	// Perform some operations
	_, err = tmpFile.WriteString("Testing manual cleanup.")

	// Check for error and cleanup
	if err != nil {
		os.Remove(tmpFile.Name())
		t.Fatal(err) // Must remember to cleanup
	}

	// Remembering to clean up at the end
	os.Remove(tmpFile.Name())
	fmt.Println("Test with manual cleanup.")
}

func main() {
	fmt.Println("Run tests with `go test`.")
}
