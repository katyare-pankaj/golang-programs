package main

import "strings"

// LowercaseProcessor implements the TextProcessor interface to convert text to lowercase
type LowercaseProcessor struct{}

// Process converts the input text to lowercase
func (l *LowercaseProcessor) Process(text string) (string, error) {
	return strings.ToLower(text), nil
}

// TextProcessor is the exported variable that must be defined in the plugin
var TextProcessor LowercaseProcessor
