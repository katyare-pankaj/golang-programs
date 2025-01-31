package main

import (
	"fmt"
	"sync"
)

// Define a callback function type
type DataProcessor func(int) int

// TransformationManager manages dynamic registration and execution of transformations.
type TransformationManager struct {
	transformations map[string]DataProcessor
	outputs         map[string]chan int
	mu              sync.Mutex
	wg              sync.WaitGroup
}

// NewTransformationManager creates a new instance of TransformationManager.
func NewTransformationManager() *TransformationManager {
	return &TransformationManager{
		transformations: make(map[string]DataProcessor),
		outputs:         make(map[string]chan int),
	}
}

// RegisterTransformation allows dynamic registration of a new transformation.
func (tm *TransformationManager) RegisterTransformation(name string, processor DataProcessor) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	tm.transformations[name] = processor
	tm.outputs[name] = make(chan int)
	tm.wg.Add(1)
	go tm.processData(name)
}

// processData manages the execution of a specific transformation.
func (tm *TransformationManager) processData(name string) {
	defer tm.wg.Done()
	for data := range tm.outputs[name] {
		result := tm.transformations[name](data)
		fmt.Printf("Transformation %s result: %d\n", name, result)
	}
}

// SendData sends data to all registered transformations.
func (tm *TransformationManager) SendData(data int) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	for _, output := range tm.outputs {
		output <- data
	}
}

// Close shuts down all running transformations and waits for completion.
func (tm *TransformationManager) Close() {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	for _, output := range tm.outputs {
		close(output)
	}
	tm.wg.Wait()
}

func main() {
	manager := NewTransformationManager()

	// Register transformations dynamically
	manager.RegisterTransformation("double", double)
	manager.RegisterTransformation("square", square)

	// Generate some sample data to process
	for i := 1; i <= 10; i++ {
		manager.SendData(i)
	}

	// Register a new transformation at runtime
	manager.RegisterTransformation("increment", increment)

	// Send more data to include the new transformation
	for i := 11; i <= 20; i++ {
		manager.SendData(i)
	}

	// Close the manager and wait for all processing to complete
	manager.Close()
	fmt.Println("Processing completed.")
}

// Sample callback function: doubles the input data
func double(data int) int {
	return data * 2
}

// Sample callback function: squares the input data
func square(data int) int {
	return data * data
}

// Sample callback function: increments the input data
func increment(data int) int {
	return data + 1
}
