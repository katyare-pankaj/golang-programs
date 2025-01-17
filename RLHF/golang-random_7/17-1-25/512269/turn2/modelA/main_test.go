package main

import (
	"errors"
	"io/ioutil"
	"os"
	"testing"
)

// FileReader is an interface for reading files, which can be mocked in tests.
type FileReader interface {
	Read(filename string) (string, error)
}

// RealFileReader implements FileReader using actual file operations.
type RealFileReader struct{}

func (r RealFileReader) Read(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// MockFileReader simulates the FileReader for testing purposes.
type MockFileReader struct {
	content string
	err     error
}

func (m MockFileReader) Read(filename string) (string, error) {
	if m.err != nil {
		return "", m.err
	}
	return m.content, nil
}

func TestReadFileSuccess(t *testing.T) {
	mockReader := MockFileReader{
		content: "mock file content",
	}

	result, err := mockReader.Read("mockfile.txt")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "mock file content"
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestReadFileError(t *testing.T) {
	mockReader := MockFileReader{
		err: errors.New("read error"),
	}

	_, err := mockReader.Read("mockfile.txt")
	if err == nil {
		t.Fatal("expected an error but got none")
	}

	expectedErrorMessage := "read error"
	if err.Error() != expectedErrorMessage {
		t.Errorf("expected error message %q, got %q", expectedErrorMessage, err.Error())
	}
}

func TestReadFileAndCleanup(t *testing.T) {
	// Create a temporary file to simulate a real file read
	tmpFile, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatalf("unexpected error creating temp file: %v", err)
	}

	// Ensure the temporary file is removed after the test
	defer os.Remove(tmpFile.Name())

	content := "temporary file content"
	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("unexpected error writing to temp file: %v", err)
	}

	// Close the file so it's ready to be read
	tmpFile.Close()

	reader := RealFileReader{}
	result, err := reader.Read(tmpFile.Name())
	if err != nil {
		t.Fatalf("unexpected error reading temp file: %v", err)
	}

	if result != content {
		t.Errorf("expected %s, got %s", content, result)
	}
}
