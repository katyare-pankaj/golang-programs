package main

import (
	"fmt"
	"sync"
	"time"
)

// CallbackFunc defines the function signature for the callbacks.
type CallbackFunc func(input, output *string)

// AIModel represents a simplified AI model capable of registering and executing callbacks.
type AIModel struct {
	callbacks []CallbackFunc
}

// NewAIModel creates and returns a new AIModel.
func NewAIModel() *AIModel {
	return &AIModel{
		callbacks: make([]CallbackFunc, 0),
	}
}

// AddCallback allows for adding a callback to be executed after prediction.
func (model *AIModel) AddCallback(callback CallbackFunc) {
	model.callbacks = append(model.callbacks, callback)
}

// Predict simulates a prediction process and executes each registered callback concurrently.
func (model *AIModel) Predict(input string) string {
	// Simulate the prediction computation.
	time.Sleep(100 * time.Millisecond) // Simulate processing delay.
	output := fmt.Sprintf("Prediction for '%s'", input)

	// Execute callbacks without waiting synchronously.
	var wg sync.WaitGroup
	wg.Add(len(model.callbacks))
	for _, callback := range model.callbacks {
		go func(cb CallbackFunc) {
			defer wg.Done()
			cb(&input, &output)
		}(callback)
	}

	wg.Wait() // Wait for all callbacks to finish.
	return output
}

func main() {
	// Initialize the model
	model := NewAIModel()

	// Logging Callback
	model.AddCallback(func(input, output *string) {
		fmt.Printf("Logging: Input='%s', Output='%s'\n", *input, *output)
	})

	// Output Modification Callback
	model.AddCallback(func(input, output *string) {
		*output = *output + " (modified)"
		fmt.Printf("Modified Output: %s\n", *output)
	})

	// Run prediction
	inputData := "example input data"
	result := model.Predict(inputData)

	// Display final output
	fmt.Printf("Final Output: %s\n", result)
}
