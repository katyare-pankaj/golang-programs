package main

import (
	"fmt"
	"sync"
	"time"
)

// Processor interface defines an extensible processor for the data pipeline.
type Processor interface {
	Process(data []byte, c chan<- []byte)
}

// SimpleProcessor simulates a basic processing function with a delay.
type SimpleProcessor struct{}

func (p *SimpleProcessor) Process(data []byte, c chan<- []byte) {
	time.Sleep(100 * time.Millisecond) // Simulate processing time
	c <- append(data, data...)
}

// AggregateProcessor simulates an aggregation processing function.
type AggregateProcessor struct{}

func (p *AggregateProcessor) Process(data []byte, c chan<- []byte) {
	time.Sleep(50 * time.Millisecond) // Simulate processing time
	c <- data
}

// Pipeline manages the sequence of processors and runs the data through them.
type Pipeline struct {
	processors []Processor
	wg         sync.WaitGroup
}

// NewPipeline creates a new pipeline.
func NewPipeline() *Pipeline {
	return &Pipeline{}
}

// AddProcessor appends a new processor to the pipeline.
func (p *Pipeline) AddProcessor(processor Processor) {
	p.processors = append(p.processors, processor)
}

// Start begins the data processing pipeline.
func (p *Pipeline) Start(input []byte) {
	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		p.runPipeline(input)
	}()
}

// Wait waits for the pipeline to complete processing.
func (p *Pipeline) Wait() {
	p.wg.Wait()
}

func (p *Pipeline) runPipeline(input []byte) {
	c := make(chan []byte)

	for _, processor := range p.processors {
		go func(p Processor) {
			processor.Process(<-c, c)
		}(processor)
	}

	c <- input
	close(c)
}

func main() {
	pipeline := NewPipeline()
	pipeline.AddProcessor(&SimpleProcessor{})
	pipeline.AddProcessor(&AggregateProcessor{})

	inputData := []byte("Hello, Fog Computing!")

	fmt.Println("Starting pipeline...")
	pipeline.Start(inputData)
	pipeline.Wait()
	fmt.Println("Pipeline completed.")
}
