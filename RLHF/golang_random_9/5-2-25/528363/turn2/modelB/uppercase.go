package main

import "strings"

// UppercaseProcessor implements the TextProcessor interface to convert text to uppercase
type UppercaseProcessor struct{}

// Process converts the input text to uppercase
func (u *UppercaseProcessor) Process(text string) (string, error) {
	return strings.ToUpper(text), nil
}

// TextProcessor is the exported variable that must be defined in the plugin
var TextProcessor UppercaseProcessor
