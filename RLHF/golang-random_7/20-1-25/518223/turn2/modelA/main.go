package main

import (
	"fmt"
	"strings"
)

// Define a function type for processing a string
type StringTransformer func(string) string

// Pipeline represents a sequence of StringTransformer operations
type Pipeline struct {
	transformers []StringTransformer
}

// Add a new transformer to the pipeline
func (p *Pipeline) Add(transformer StringTransformer) {
	p.transformers = append(p.transformers, transformer)
}

// Execute runs all transformers in sequence on the input string
func (p *Pipeline) Execute(input string) string {
	current := input
	for _, transformer := range p.transformers {
		current = transformer(current)
	}
	return current
}

// Transformer functions
func UpperCaseTransformer(s string) string {
	return strings.ToUpper(s)
}

func ReverseTransformer(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func CustomFormatTransformer(s string) string {
	return fmt.Sprintf("[%s]", s)
}

func main() {
	// Create a new processing pipeline
	pipeline := &Pipeline{}

	// Add transformers to the pipeline
	pipeline.Add(UpperCaseTransformer)
	pipeline.Add(ReverseTransformer)
	pipeline.Add(CustomFormatTransformer)

	// Sample input string
	input := "hello, world"

	// Execute the pipeline and obtain the transformed result
	transformed := pipeline.Execute(input)

	fmt.Println("Original String:", input)
	fmt.Println("Transformed String:", transformed)

	// Example of a different pipeline with a different order of operations
	anotherPipeline := &Pipeline{}
	anotherPipeline.Add(CustomFormatTransformer)
	anotherPipeline.Add(UpperCaseTransformer)
	anotherPipeline.Add(ReverseTransformer)

	transformedAgain := anotherPipeline.Execute(input)
	fmt.Println("Transformed Again:", transformedAgain)
}
