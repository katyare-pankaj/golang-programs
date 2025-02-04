package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// FileManager defines an interface for file operations
type FileManager interface {
	Read(filename string) (string, error)
	Write(filename string, data string) error
	Delete(filename string) error
}

// DiskFileManager implements the FileManager interface for disk storage
type DiskFileManager struct{}

func (dfm *DiskFileManager) Read(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (dfm *DiskFileManager) Write(filename string, data string) error {
	return ioutil.WriteFile(filename, []byte(data), 0644)
}

func (dfm *DiskFileManager) Delete(filename string) error {
	return os.Remove(filename)
}

// InMemoryFileManager implements the FileManager interface for in-memory storage
type InMemoryFileManager struct {
	files map[string]string
}

func NewInMemoryFileManager() *InMemoryFileManager {
	return &InMemoryFileManager{files: make(map[string]string)}
}

func (imfm *InMemoryFileManager) Read(filename string) (string, error) {
	data, exists := imfm.files[filename]
	if !exists {
		return "", fmt.Errorf("file not found: %s", filename)
	}
	return data, nil
}

func (imfm *InMemoryFileManager) Write(filename string, data string) error {
	imfm.files[filename] = data
	return nil
}

func (imfm *InMemoryFileManager) Delete(filename string) error {
	delete(imfm.files, filename)
	return nil
}

// DemonstrateFileOperations shows how to use the FileManager interface
func DemonstrateFileOperations(fm FileManager, filename string, content string) error {
	fmt.Printf("Writing to %s...\n", filename)
	err := fm.Write(filename, content)
	if err != nil {
		return fmt.Errorf("write error: %w", err)
	}

	fmt.Printf("Reading from %s...\n", filename)
	readContent, err := fm.Read(filename)
	if err != nil {
		return fmt.Errorf("read error: %w", err)
	}
	fmt.Printf("Content read: %s\n", readContent)

	fmt.Printf("Deleting %s...\n", filename)
	err = fm.Delete(filename)
	if err != nil {
		return fmt.Errorf("delete error: %w", err)
	}

	return nil
}

func main() {
	// Using DiskFileManager
	fmt.Println("Disk File Manager:")
	diskManager := &DiskFileManager{}
	if err := DemonstrateFileOperations(diskManager, "sample.txt", "Hello, Disk!"); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}

	fmt.Println()

	// Using InMemoryFileManager
	fmt.Println("In-Memory File Manager:")
	memManager := NewInMemoryFileManager()
	if err := DemonstrateFileOperations(memManager, "sample.txt", "Hello, Memory!"); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
}
