// readplugin.go
package main

import (
	"fmt"
	"io/ioutil"
)

// ReadOperation implements the FileOperation interface to read files
type ReadOperation struct{}

// Perform reads the content of the specified file
func (r *ReadOperation) Perform(filePath string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	fmt.Println("File content:", string(data))
	return nil
}

// FileOperation implements the Plugin's FileOperation interface
var FileOperation ReadOperation
