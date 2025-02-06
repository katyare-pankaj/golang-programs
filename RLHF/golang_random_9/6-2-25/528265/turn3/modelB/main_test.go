package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestTempFile(t *testing.T) {
	// Acquire a temporary file
	file, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatal(err)
	}

	// Close and remove the file at the end of the test
	defer func() {
		file.Close()
		os.Remove(file.Name())
	}()

	// Use the file for testing
	_, err = file.WriteString("This is a test file.")
	if err != nil {
		t.Fatal(err)
	}

	// Add more test assertions here...
}
