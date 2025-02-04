package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

// Lifecycle represents the different states of a file's lifecycle.
type Lifecycle interface {
	Process(string) error
}

// ActiveLifecycle represents the active state of a file's lifecycle.
type ActiveLifecycle struct{}

// Process processes the active lifecycle of a file.
func (a *ActiveLifecycle) Process(filePath string) error {
	fmt.Printf("Processing active file: %s\n", filePath)
	// Add your custom active file processing logic here
	return nil
}

// InactiveLifecycle represents the inactive state of a file's lifecycle.
type InactiveLifecycle struct{}

// Process processes the inactive lifecycle of a file.
func (i *InactiveLifecycle) Process(filePath string) error {
	fmt.Printf("Processing inactive file: %s\n", filePath)
	// Add your custom inactive file processing logic here
	return nil
}

// FileManager manages the lifecycle of files based on their state.
type FileManager struct {
	lifecycle Lifecycle
}

// NewFileManager creates a new FileManager with the specified lifecycle.
func NewFileManager(lifecycle Lifecycle) *FileManager {
	return &FileManager{lifecycle: lifecycle}
}

// ProcessFiles processes all the files in the specified directory with the specified lifecycle.
func (fm *FileManager) ProcessFiles(dirPath string) error {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("error reading directory %s: %w", dirPath, err)
	}

	for _, file := range files {
		filePath := filepath.Join(dirPath, file.Name())
		if file.IsDir() {
			err := fm.ProcessFiles(filePath)
			if err != nil {
				return fmt.Errorf("error processing subdirectory %s: %w", filePath, err)
			}
		} else {
			err := fm.lifecycle.Process(filePath)
			if err != nil {
				return fmt.Errorf("error processing file %s: %w", filePath, err)
			}
		}
	}
	return nil
}

func main() {
	// Create a new FileManager with an ActiveLifecycle
	activeFileManager := NewFileManager(&ActiveLifecycle{})

	// Create a new FileManager with an InactiveLifecycle
	inactiveFileManager := NewFileManager(&InactiveLifecycle{})

	dataDirPath := "./data" // Replace this with your actual data directory path

	// Process active files
	err := activeFileManager.ProcessFiles(dataDirPath)
	if err != nil {
		log.Fatalf("Error processing active files: %s", err)
	}

	// Process inactive files
	err = inactiveFileManager.ProcessFiles(dataDirPath)
	if err != nil {
		log.Fatalf("Error processing inactive files: %s", err)
	}
}
