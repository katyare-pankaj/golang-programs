package main

import (
	"fmt"
	"sync"
	"time"
)

// CallbackFunc defines the function signature for callbacks.
type CallbackFunc func(input, output *string)

// AIModel represents an AI model capable of handling concurrent callbacks.
type AIModel struct {
	callbacks []CallbackFunc
}

// NewAIModel initializes and returns a new AIModel instance.
func NewAIModel() *AIModel {
	return &AIModel{
		callbacks: make([]CallbackFunc, 0),
	}
}

// AddCallback adds a callback to the AI model.
func (model *AIModel) AddCallback(callback CallbackFunc) {
	model.callbacks = append(model.callbacks, callback)
}

// Predict simulates a prediction and executes registered callbacks concurrently.
func (model *AIModel) Predict(input string) string {
	// Simulate prediction processing
	time.Sleep(50 * time.Millisecond) // Simulate a quick processing delay
	output := fmt.Sprintf("Prediction for '%s'", input)

	// Initialize a wait group for concurrent callback execution
	var wg sync.WaitGroup
	wg.Add(len(model.callbacks))

	// Execute each callback in a separate goroutine
	for _, callback := range model.callbacks {
		go func(cb CallbackFunc) {
			defer wg.Done()
			cb(&input, &output)
		}(callback)
	}

	// Return the prediction immediately
	go func() {
		wg.Wait() // Wait for all callbacks to complete
	}()

	return output
}

func main() {
	// Create and configure the AI model
	model := NewAIModel()

	// Example logging callback
	model.AddCallback(func(input, output *string) {
		fmt.Printf("Logging: Input='%s', Output='%s'\n", *input, *output)
	})

	// Example modification callback
	model.AddCallback(func(input, output *string) {
		*output = *output + " [processed]"
		fmt.Printf("Modified Output: %s\n", *output)
	})

	// Run a prediction and obtain the result
	inputData := "example input"
	result := model.Predict(inputData)

	// Output the prediction result
	fmt.Printf("Initial Output: %s\n", result)

	// Allow some time for callbacks to complete before finish
	time.Sleep(100 * time.Millisecond)
}
