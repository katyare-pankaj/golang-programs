package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// DataProcessor represents a generic data processor in the pipeline
type DataProcessor interface {
	Process(ctx context.Context, data []byte) []byte
}

// SimpleProcessor is an example of a basic data processor
type SimpleProcessor struct {
}

func (p *SimpleProcessor) Process(ctx context.Context, data []byte) []byte {
	// Simulate some basic data processing
	time.Sleep(10 * time.Millisecond)
	return append(data, []byte(" processed")...)
}

// Pipeline manages the data processing pipeline
type Pipeline struct {
	processors []DataProcessor
	wg         sync.WaitGroup
}

func NewPipeline(processors ...DataProcessor) *Pipeline {
	return &Pipeline{processors: processors}
}

func (p *Pipeline) Start(ctx context.Context, data []byte) {
	p.wg.Add(1)
	defer p.wg.Done()

	go func() {
		processedData := data
		for _, processor := range p.processors {
			processedData = processor.Process(ctx, processedData)
		}
		fmt.Println("Final processed data:", string(processedData))
	}()
}

func main() {
	ctx := context.Background()
	data := []byte("Hello, Fog Computing!")

	// Create an extensible pipeline with multiple processors
	pipeline := NewPipeline(&SimpleProcessor{}, &SimpleProcessor{}, &SimpleProcessor{})

	pipeline.Start(ctx, data)
	pipeline.wg.Wait()
}
