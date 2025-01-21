package main

import (
	"fmt"
	"strings"
)

// Define a function type for processing strings
type StringTransformer func(string) string

// Pipeline represents a sequence of StringTransformer operations
type Pipeline struct {
	transformers []StringTransformer
}

// Add appends a new transformer to the pipeline
func (p *Pipeline) Add(transformer StringTransformer) {
	p.transformers = append(p.transformers, transformer)
}

// Remove deletes a transformer from the pipeline by index
func (p *Pipeline) Remove(index int) error {
	if index < 0 || index >= len(p.transformers) {
		return fmt.Errorf("index out of range: %d", index)
	}
	p.transformers = append(p.transformers[:index], p.transformers[index+1:]...)
	return nil
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

	fmt.Println("Pipeline with all steps:")
	input := "hello, world"
	fmt.Println("Original String:", input)
	fmt.Println("Transformed String:", pipeline.Execute(input))

	// Remove the second transformer (ReverseTransformer)
	err := pipeline.Remove(1)
	if err != nil {
		fmt.Println("Error removing transformer:", err)
	}

	fmt.Println("\nPipeline after removing a step:")
	fmt.Println("Transformed String:", pipeline.Execute(input))

	// Users can dynamically add their own transformation
	pipeline.Add(func(s string) string { return s + "!!!" })

	fmt.Println("\nPipeline after adding a new step:")
	fmt.Println("Transformed String:", pipeline.Execute(input))
}
