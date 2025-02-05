package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Implement the FileProcessor interface
type uppercaseProcessor struct{}

func (up *uppercaseProcessor) Process(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	upperData := strings.ToUpper(string(data))
	err = ioutil.WriteFile(filename, []byte(upperData), 0644)
	if err != nil {
		return err
	}
	fmt.Println("Converted file to uppercase.")
	return nil
}

// Expose the plugin symbol
var FileProcessor FileProcessor = &uppercaseProcessor{}
