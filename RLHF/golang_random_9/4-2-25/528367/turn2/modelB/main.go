package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"
)

// FileManager defines the contract for file operations
type FileManager interface {
	Read(filename string) ([]byte, error)
	Write(filename string, data []byte) error
	Delete(filename string) error
	Exists(filename string) bool
}

// DiskFileManager implements the FileManager for disk storage
type DiskFileManager struct {
	baseDir string
}

// NewDiskFileManager creates a new DiskFileManager with a given base directory
func NewDiskFileManager(baseDir string) *DiskFileManager {
	return &DiskFileManager{baseDir: baseDir}
}

// Read reads data from a file on disk
func (dfm *DiskFileManager) Read(filename string) ([]byte, error) {
	path := dfm.getFilePath(filename)
	return ioutil.ReadFile(path)
}

// Write writes data to a file on disk
func (dfm *DiskFileManager) Write(filename string, data []byte) error {
	path := dfm.getFilePath(filename)
	return ioutil.WriteFile(path, data, 0644)
}

// Delete removes a file from disk
func (dfm *DiskFileManager) Delete(filename string) error {
	path := dfm.getFilePath(filename)
	return os.Remove(path)
}

// Exists checks if a file exists on disk
func (dfm *DiskFileManager) Exists(filename string) bool {
	path := dfm.getFilePath(filename)
	_, err := os.Stat(path)
	return err == nil
}

// getFilePath generates the full path for a given filename
func (dfm *DiskFileManager) getFilePath(filename string) string {
	return fmt.Sprintf("%s/%s", dfm.baseDir, filename)
}

// InMemoryFileManager implements the FileManager for in-memory storage, useful for testing
type InMemoryFileManager struct {
	files   map[string][]byte
	mu      sync.RWMutex
}

// NewInMemoryFileManager initializes an InMemoryFileManager
func NewInMemoryFileManager() *InMemoryFileManager {
	return &InMemoryFileManager{files: make(map[string][]byte)}
}

// Read reads data from an in-memory file
func (imfm *InMemoryFileManager) Read(filename string) ([]byte, error) {
	imfm.mu.RLock()
	defer imfm.mu.RUnlock()
	if data, exists := imfm.files[filename]; exists {
		return data, nil
	}
	return nil, fmt.Errorf("file not found: %s", filename)
}

// Write writes data to an in-memory file
func (imfm *InMemoryFileManager) Write(filename string, data []byte) error {
	imfm.mu.Lock()
	defer imfm.mu.Unlock()
	imfm.files[filename] = data
	return nil
}

// Delete removes an in-memory file
func (imfm *InMemoryFileManager) Delete(filename string) error {
	imfm.mu.Lock()
	defer imfm.mu.Unlock()
	delete(imfm.files, filename)
	return nil
}

// Exists checks if an in-memory file exists
func (imfm *InMemoryFileManager) Exists(filename string) bool {
	imfm.mu.RLock()
	defer imfm.mu.RUnlock()
	_, exists := imfm.files[filename]
	return exists
}

// ProcessFile demonstrates using the FileManager interface to manipulate a file
func ProcessFile(fm FileManager, filename string, content string) error {
	// Write to the file
	err := fm.Write(filename, []byte(content))
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}