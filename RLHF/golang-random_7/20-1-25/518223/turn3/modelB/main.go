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

// Add a new transformer to the back of the pipeline
func (p *Pipeline) Add(transformer StringTransformer) {
	p.transformers = append(p.transformers, transformer)
}

// AddTransformerAt dynamically adds a transformer at the specified index.
// It shifts existing transformers after the index right.
func (p *Pipeline) AddTransformerAt(index int, transformer StringTransformer) {
	if index < 0 || index > len(p.transformers) {
		fmt.Println("Invalid index to add transformer. Index must be between 0 and", len(p.transformers))
		return
	}
	p.transformers = append(p.transformers[:index], append([]StringTransformer{transformer}, p.transformers[index:]...)...)
}

// RemoveTransformerAt dynamically removes a transformer at the specified index.
// It shifts remaining transformers left.
func (p *Pipeline) RemoveTransformerAt(index int) {
	if index < 0 || index >= len(p.transformers) {
		fmt.Println("Invalid index to remove transformer. Index must be between 0 and", len(p.transformers)-1)
		return
	}
	p.transformers = append(p.transformers[:index], p.transformers[index+1:]...)
}

// Execute runs all transformers in sequence on the input string
func (p *Pipeline) Execute(input string) string {
	current := input
	for _, transformer := range p.transformers {
		current = transformer(current)
	}
	return current
}

// Transformer functions (the same as before)
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
	pipeline := &Pipeline{}

	// Dynamically add transformers at runtime
	pipeline.AddTransformerAt(0, UpperCaseTransformer)
	pipeline.AddTransformerAt(1, ReverseTransformer)
	pipeline.AddTransformerAt(2, CustomFormatTransformer)

	fmt.Println("Pipeline:", pipeline.transformers)

	input := "hello, world"
	transformed := pipeline.Execute(input)
	fmt.Println("Transformed String:", transformed)

	// Dynamically remove transformer at index 1 (ReverseTransformer)
	pipeline.RemoveTransformerAt(1)
	transformed = pipeline.Execute(input)
	fmt.Println("Transformed String after removing ReverseTransformer:", transformed)
}
