package main

import (
	"fmt"
	"sync"
)

// Processor interface defines an extensible data processor.
type Processor interface {
	Process(in <-chan []byte, out chan<- []byte)
}

// SimpleProcessor doubles the data.
type SimpleProcessor struct{}

func (p *SimpleProcessor) Process(in <-chan []byte, out chan<- []byte) {
	for data := range in {
		out <- append(data, data...)
	}
}

// AggregateProcessor aggregates the data.
type AggregateProcessor struct{}

func (p *AggregateProcessor) Process(in <-chan []byte, out chan<- []byte) {
	for data := range in {
		// Simulate aggregation
		out <- data
	}
}

// Pipeline runs a sequence of processors.
type Pipeline struct {
	processors []Processor
	wg         sync.WaitGroup
}

// NewPipeline creates a new pipeline.
func NewPipeline(processors ...Processor) *Pipeline {
	return &Pipeline{processors: processors}
}

// Start runs the pipeline.
func (p *Pipeline) Start() {
	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		in := make(chan []byte)
		out := make(chan []byte)

		for _, processor := range p.processors {
			go processor.Process(in, out)
			in = out
			out = make(chan []byte)
		}

		// Simulate data input
		for i := 0; i < 1000; i++ {
			data := []byte(fmt.Sprintf("Data %d\n", i))
			in <- data
		}

		close(in)
	}()
}

// Wait waits for the pipeline to finish.
func (p *Pipeline) Wait() {
	p.wg.Wait()
}

func main() {
	pipeline := NewPipeline(&SimpleProcessor{}, &AggregateProcessor{})
	pipeline.Start()
	pipeline.Wait()

	fmt.Println("Pipeline processing complete.")
}
