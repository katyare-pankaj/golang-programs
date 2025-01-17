package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

type FileReader interface {
	Read(p []byte) (n int, err error)
}

func processFile(reader FileReader) (string, error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

type MockFileReader struct {
	Data []byte
	Err  error
}

func (r *MockFileReader) Read(p []byte) (n int, err error) {
	if r.Err != nil {
		return 0, r.Err
	}
	n = copy(p, r.Data)
	return n, io.EOF
}

func TestProcessFile(t *testing.T) {
	// Create a temporary file with known content
	tmpFile, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatalf("failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up the temporary file

	_, err = tmpFile.WriteString("hello, world!")
	if err != nil {
		t.Fatalf("failed to write to temporary file: %v", err)
	}
	if err := tmpFile.Close(); err != nil {
		t.Fatalf("failed to close temporary file: %v", err)
	}

	// Test successful processing
	mockReader := &MockFileReader{Data: []byte("hello, world!")}
	result, err := processFile(mockReader)
	if err != nil {
		t.Errorf("unexpected error processing file: %v", err)
	}
	expected := "hello, world!"
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}

	// Test read error
	mockReaderWithError := &MockFileReader{Err: fmt.Errorf("read error")}
	_, err = processFile(mockReaderWithError)
	if err == nil {
		t.Errorf("expected read error, got nil")
	}
	if err.Error() != "read error" {
		t.Errorf("expected error message 'read error', got %q", err.Error())
	}
}
