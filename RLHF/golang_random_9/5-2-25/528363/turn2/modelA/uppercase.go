// uppercase.go
package main

import (
	"errors"
	"strings"
)

// UppercaseProcessor implements the TextProcessor interface
type UppercaseProcessor struct{}

// Process converts the input text to uppercase
func (u *UppercaseProcessor) Process(text string) (string, error) {
	if text == "" {
		return "", errors.New("input text cannot be empty")
	}
	return strings.ToUpper(text), nil
}

// Export the symbol for the plugin
var TextProcessorPlugin UppercaseProcessor
