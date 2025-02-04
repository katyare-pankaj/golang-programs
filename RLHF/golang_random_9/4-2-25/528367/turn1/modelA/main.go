package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// FileManager defines the contract for file operations
type FileManager interface {
	Read(filename string) (string, error)
	Write(filename string, data string) error
	Delete(filename string) error
}

// DiskFileManager implements the FileManager for disk storage
type DiskFileManager struct{}

// Read reads data from a file on disk
func (dfm *DiskFileManager) Read(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// Write writes data to a file on disk
func (dfm *DiskFileManager) Write(filename string, data string) error {
	return ioutil.WriteFile(filename, []byte(data), 0644)
}

// Delete removes a file from disk
func (dfm *DiskFileManager) Delete(filename string) error {
	return os.Remove(filename)
}

// InMemoryFileManager implements the FileManager for in-memory storage, useful for testing
type InMemoryFileManager struct {
	files map[string]string
}

// NewInMemoryFileManager initializes an InMemoryFileManager
func NewInMemoryFileManager() *InMemoryFileManager {
	return &InMemoryFileManager{files: make(map[string]string)}
}

// Read reads data from an in-memory file
func (imfm *InMemoryFileManager) Read(filename string) (string, error) {
	if data, exists := imfm.files[filename]; exists {
		return data, nil
	}
	return "", fmt.Errorf("file not found: %s", filename)
}

// Write writes data to an in-memory file
func (imfm *InMemoryFileManager) Write(filename string, data string) error {
	imfm.files[filename] = data
	return nil
}

// Delete removes an in-memory file
func (imfm *InMemoryFileManager) Delete(filename string) error {
	delete(imfm.files, filename)
	return nil
}

// ProcessFile demonstrates using the FileManager interface to manipulate a file
func ProcessFile(fm FileManager, filename string, content string) error {
	// Write to the file
	err := fm.Write(filename, content)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	// Read from the file
	readContent, err := fm.Read(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}
	fmt.Printf("Read content: %s\n", readContent)

	// Delete the file
	err = fm.Delete(filename)
	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}

func main() {
	fmt.Println("Disk File Manager:")
	diskFM := &DiskFileManager{}
	err := ProcessFile(diskFM, "sample.txt", "Hello, World!")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}

	fmt.Println("\nIn-Memory File Manager:")
	memFM := NewInMemoryFileManager()
	err = ProcessFile(memFM, "sample.txt", "Hello, World!")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
}
