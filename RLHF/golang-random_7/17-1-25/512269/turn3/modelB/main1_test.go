package main

import (
	"errors"
	"testing"
)

// MockFileReader simulates a file reader
type MockFileReader struct {
	Data []byte
}

// Read is a method that mimics reading from a file.
func (m *MockFileReader) Read(p []byte) (n int, err error) {
	copy(p, m.Data)
	return len(m.Data), nil
}

// processFile processes the file data and returns it as a string.
func processFile(reader *MockFileReader) (string, error) {
	if reader == nil {
		return "", errors.New("reader is nil")
	}
	buf := make([]byte, len(reader.Data))
	_, err := reader.Read(buf)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

// TestProcessFile_EmptyFile tests an empty file scenario.
func TestProcessFile_EmptyFile(t *testing.T) {
	mockReader := &MockFileReader{Data: []byte{}}
	result, err := processFile(mockReader)
	if err != nil {
		t.Errorf("unexpected error processing file: %v", err)
	}
	expected := ""
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

// TestProcessFile_LargeFile tests a large file scenario.
func TestProcessFile_LargeFile(t *testing.T) {
	const largeDataSize = 1024 * 1024 // 1 MB
	largeData := make([]byte, largeDataSize)
	mockReader := &MockFileReader{Data: largeData}
	result, err := processFile(mockReader)
	if err != nil {
		t.Errorf("unexpected error processing file: %v", err)
	}
	expected := string(largeData)
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

// TestProcessFile_BinaryFile tests a binary file scenario.
func TestProcessFile_BinaryFile(t *testing.T) {
	binaryData := []byte{0x00, 0x01, 0x02, 0x03, 0x04}
	mockReader := &MockFileReader{Data: binaryData}
	result, err := processFile(mockReader)
	if err != nil {
		t.Errorf("unexpected error processing file: %v", err)
	}
	expected := string(binaryData)
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}
