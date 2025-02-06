package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestManualCleanup(t *testing.T) {
	// Create a temporary file
	tmpFile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Fatal(err)
	}
	cleanup := true
	manualCleanup := func() {
		if cleanup {
			os.Remove(tmpFile.Name())
		}
	}

	// Ensure that manualCleanup is only called once
	defer func() {
		if !t.Failed() {
			manualCleanup()
		}
	}()

	// Perform test operations
	_, err = tmpFile.WriteString("Hello, Gophers!")
	if err != nil {
		t.Fatal(err)
	}

	// A condition where we may decide not to clean up
	if someCondition() {
		cleanup = false
	}

	// Additional test logic...
}

func someCondition() bool {
	// Simulates some condition - returning false to ensure cleanup
	return false
}
