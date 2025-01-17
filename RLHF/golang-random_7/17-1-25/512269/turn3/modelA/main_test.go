package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

// FileReader interface
type FileReader interface {
	Read(p []byte) (n int, err error)
}

// Function to process the file
func processFile(reader FileReader) (string, error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// MockFileReader struct
type MockFileReader struct {
	Data []byte
	Err  error
}

// Implementing the Read method
func (r *MockFileReader) Read(p []byte) (n int, err error) {
	if r.Err != nil {
		return 0, r.Err
	}
	n = copy(p, r.Data)
	// Simulate reaching the end of the file
	if n < len(r.Data) {
		return n, nil
	}
	return n, io.EOF
}

func TestProcessFile(t *testing.T) {
	tests := []struct {
		name        string
		reader      *MockFileReader
		expected    string
		expectError bool
	}{
		{
			name:        "Empty file",
			reader:      &MockFileReader{Data: []byte("")},
			expected:    "",
			expectError: false,
		},
		{
			name:        "Small text file",
			reader:      &MockFileReader{Data: []byte("hello, world!")},
			expected:    "hello, world!",
			expectError: false,
		},
		{
			name:        "Large text file",
			reader:      &MockFileReader{Data: []byte(strings.Repeat("a", 1024*1024))}, // 1 MB of 'a'
			expected:    strings.Repeat("a", 1024*1024),
			expectError: false,
		},
		{
			name:        "Binary file",
			reader:      &MockFileReader{Data: []byte{0x00, 0xFF, 0xAA, 0xBB}},
			expected:    string([]byte{0x00, 0xFF, 0xAA, 0xBB}),
			expectError: false,
		},
		{
			name:        "Read error",
			reader:      &MockFileReader{Err: fmt.Errorf("read error")},
			expected:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := processFile(tt.reader)
			if tt.expectError {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("expected %q, got %q", tt.expected, result)
				}
			}
		})
	}
}

func main() {
	fmt.Println("Run the tests using `go test` command.")
}
