package main

import (
	"testing"
)

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
