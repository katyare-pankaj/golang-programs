package main

import "fmt"

// Processor interface defines an extensible processor for the data pipeline.
type Processor interface {
	Process(input []byte) []byte
}

// SimpleProcessor implements a basic processing function.
type SimpleProcessor struct{}

func (p *SimpleProcessor) Process(input []byte) []byte {
	// Simulate some basic processing
	return append(input, input...)
}

// AggregateProcessor implements an aggregation processing function.
type AggregateProcessor struct{}

func (p *AggregateProcessor) Process(input []byte) []byte {
	// Simulate aggregation logic
	// For demonstration purposes, we simply return the original data
	return input
}

// Pipeline encapsulates a sequence of processors.
type Pipeline struct {
	processors []Processor
}

// NewPipeline creates a new pipeline with an initial processor.
func NewPipeline(first Processor) *Pipeline {
	return &Pipeline{processors: []Processor{first}}
}

// Append adds a new processor to the end of the pipeline.
func (p *Pipeline) Append(processor Processor) {
	p.processors = append(p.processors, processor)
}

// Process runs the data through the entire pipeline.
func (p *Pipeline) Process(input []byte) []byte {
	output := input
	for _, processor := range p.processors {
		output = processor.Process(output)
	}
	return output
}

// Example usage:
func main() {
	pipeline := NewPipeline(&SimpleProcessor{})
	pipeline.Append(&AggregateProcessor{})

	inputData := []byte("Hello, Fog Computing!")
	result := pipeline.Process(inputData)

	fmt.Println("Result:", string(result))
}
