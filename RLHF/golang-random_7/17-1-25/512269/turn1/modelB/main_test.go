package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type reader interface {
	Read() ([]byte, error)
}

type fileReader struct {
	filename string
}

func (fr fileReader) Read() ([]byte, error) {
	return ioutil.ReadFile(fr.filename)
}

func processFile(r reader) error {
	// Read the content from the reader
	content, err := r.Read()
	if err != nil {
		return err
	}

	// Do some processing with the content
	fmt.Println(string(content))
	return nil
}

func TestProcessFileWithDefer(t *testing.T) {
	t.Run("Should successfully process the file", func(t *testing.T) {
		// Create a temporary file and write some content
		tempFile, err := ioutil.TempFile("", "testfile")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tempFile.Name()) // Clean up the temporary file after the test

		testContent := []byte("Hello, world!")
		if _, err := tempFile.Write(testContent); err != nil {
			t.Fatal(err)
		}

		// Create a mock reader that returns the content of the temporary file
		mockReader := &mockReader{
			readBytes: testContent,
			readErr:   nil,
		}

		// Process the file using the mock reader
		if err := processFile(mockReader); err != nil {
			t.Fatal(err)
		}

		// Verify that the content of the mock reader was printed correctly
		expectedOutput := string(testContent)
		if actualOutput := getStdoutOutput(); actualOutput != expectedOutput {
			t.Fatalf("Expected output: %s, Actual output: %s", expectedOutput, actualOutput)
		}
	})

	t.Run("Should handle read error", func(t *testing.T) {
		// Create a mock reader that returns an error during read
		mockReader := &mockReader{
			readBytes: nil,
			readErr:   errors.New("some read error"),
		}

		// Process the file using the mock reader
		err := processFile(mockReader)
		if err == nil {
			t.Fatal("Expected an error, but got nil")
		}

		// Verify that the correct error message is returned
		expectedErrorMessage := "some read error"
		if actualErrorMessage := err.Error(); actualErrorMessage != expectedErrorMessage {
			t.Fatalf("Expected error message: %s, Actual error message: %s", expectedErrorMessage, actualErrorMessage)
		}
	})
}

// mockReader is a mock object that implements the reader interface.
// It can be used to simulate the behavior of a real file reader.
type mockReader struct {
	readBytes []byte
	readErr   error
}

func (mr *mockReader) Read() ([]byte, error) {
	return mr.readBytes, mr.readErr
}
