package main

import (
	"fmt"
	"log"
	"os"
	"plugin"
)

// Define the interface that plugins must implement
type FileProcessor interface {
	Process(filename string) error
}

// Main function
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <plugin_name>")
		return
	}
	pluginName := os.Args[1]

	// Load the plugin
	p, err := plugin.Open(pluginName + ".so")
	if err != nil {
		log.Fatalf("Error loading plugin: %v", err)
	}

	// Look up the symbol for the FileProcessor interface
	symFileProcessor, err := p.Lookup("FileProcessor")
	if err != nil {
		log.Fatalf("Error looking up symbol: %v", err)
	}

	// Assert that the symbol implements the FileProcessor interface
	var fileProcessor FileProcessor
	if fp, ok := symFileProcessor.(FileProcessor); ok {
		fileProcessor = fp
	} else {
		log.Fatalf("Symbol does not implement FileProcessor interface")
	}

	// Run the plugin's Process function
	filename := "data.txt"
	if len(os.Args) > 2 {
		filename = os.Args[2]
	}
	err = fileProcessor.Process(filename)
	if err != nil {
		log.Fatalf("Error processing file: %v", err)
	}
	fmt.Println("File processing completed successfully.")
}
