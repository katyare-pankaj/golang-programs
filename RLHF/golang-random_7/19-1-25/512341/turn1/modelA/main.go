package main

import (
	"fmt"
	"time"
)

// CallbackFunc defines the type for callback functions.
type CallbackFunc func(input, output interface{})

// SimpleAIModel represents a simple AI model with callback support.
type SimpleAIModel struct {
	// Add model fields here (e.g., weights, biases)
	callbacks []CallbackFunc
}

// NewSimpleAIModel initializes a new SimpleAIModel.
func NewSimpleAIModel() *SimpleAIModel {
	return &SimpleAIModel{
		callbacks: make([]CallbackFunc, 0),
	}
}

// AddCallback adds a callback function to the model.
func (model *SimpleAIModel) AddCallback(callback CallbackFunc) {
	model.callbacks = append(model.callbacks, callback)
}

// Predict simulates a prediction function and triggers callbacks.
func (model *SimpleAIModel) Predict(input interface{}) interface{} {
	// Simulate prediction processing.
	time.Sleep(100 * time.Millisecond) // Simulate some processing time.
	output := fmt.Sprintf("prediction for %v", input)

	// Execute callbacks concurrently to maintain performance.
	done := make(chan struct{})
	for _, callback := range model.callbacks {
		go func(cb CallbackFunc) {
			cb(input, output)
			done <- struct{}{}
		}(callback)
	}

	// Wait for all callbacks to complete.
	for i := 0; i < len(model.callbacks); i++ {
		<-done
	}

	return output
}

func main() {
	// Initialize the AI model and add callbacks.
	model := NewSimpleAIModel()

	// Example callback for logging
	model.AddCallback(func(input, output interface{}) {
		fmt.Printf("Logging: input=%v, output=%v\n", input, output)
	})

	// Example callback for modifying output
	model.AddCallback(func(input, output interface{}) {
		if out, ok := output.(string); ok {
			modifiedOutput := out + " (modified)"
			fmt.Printf("Modified Output: %s\n", modifiedOutput)
		}
	})

	// Execute a prediction.
	inputData := "example input"
	result := model.Predict(inputData)

	// Output prediction result.
	fmt.Printf("Final result: %v\n", result)
}
