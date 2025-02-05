// core.go
package main

import (
	"fmt"
	"log"
	"os"
	"plugin"
)

// FileOperation defines the interface for plugins to implement
type FileOperation interface {
	Perform(filePath string) error
}

// LoadPlugin loads the specified plugin and returns a FileOperation implementation
func LoadPlugin(pluginPath string) (FileOperation, error) {
	p, err := plugin.Open(pluginPath)
	if err != nil {
		return nil, err
	}

	symFileOperation, err := p.Lookup("FileOperation")
	if err != nil {
		return nil, err
	}

	var fileOp FileOperation
	fileOp, ok := symFileOperation.(FileOperation)
	if !ok {
		return nil, fmt.Errorf("unexpected type from module symbol")
	}

	return fileOp, nil
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <plugin-path> <file-path>\n", os.Args[0])
	}

	pluginPath := os.Args[1]
	filePath := os.Args[2]

	fileOp, err := LoadPlugin(pluginPath)
	if err != nil {
		log.Fatalf("Could not load plugin: %v", err)
	}

	if err := fileOp.Perform(filePath); err != nil {
		log.Fatalf("Error performing operation: %v", err)
	}

	fmt.Println("Operation completed successfully")
}
