package main

import (
	"fmt"
	"time"
)

// CallbackFunc defines the type for callback functions.
type CallbackFunc func(input, output interface{})

// SimpleAIModel represents a simple AI model with callback support.
type SimpleAIModel struct {
	// Add model fields here (e.g., weights, biases).
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

	return output
}

// StartPredictionLoop starts a loop to predict inputs and trigger callbacks.
func (model *SimpleAIModel) StartPredictionLoop(inputs <-chan interface{}, callbackDone chan<- struct{}) {
	callbackData := make(chan struct{ input, output interface{} })
	go func() {
		for {
			select {
			case input := <-inputs:
				output := model.Predict(input)
				callbackData <- struct{ input, output interface{} }{input, output}
			case <-callbackDone:
				return
			}
		}
	}()

	// Spawn goroutines to handle callbacks
	numCallbackWorkers := len(model.callbacks)
	callbackWorkers := make(chan struct{}, numCallbackWorkers)
	for _, callback := range model.callbacks {
		go func(cb CallbackFunc) {
			for {
				select {
				case data := <-callbackData:
					cb(data.input, data.output)
					callbackWorkers <- struct{}{}
				case <-callbackDone:
					return
				}
			}
		}(callback)
	}

	// Wait for all callback workers to finish.
	for i := 0; i < numCallbackWorkers; i++ {
		<-callbackWorkers
	}
}

func main() {
	// Initialize the AI model and add callbacks.
	model := NewSimpleAIModel()
	model.AddCallback(func(input, output interface{}) {
		fmt.Printf("Logging: input=%v, output=%v\n", input, output)
	})

	model.AddCallback(func(input, output interface{}) {
		if out, ok := output.(string); ok {
			modifiedOutput := out + " (modified)"
			fmt.Printf("Modified Output: %s\n", modifiedOutput)
		}
	})

	// Create input channel and start the prediction loop
	inputs := make(chan interface{})
	callbackDone := make(chan struct{})
	go model.StartPredictionLoop(inputs, callbackDone)

	// Simulate sending inputs to the prediction loop
	inputData := []interface{}{"example1", "example2", "example3"}
	for _, input := range inputData {
		inputs <- input
	}

	// Close the input channel to signal the end of predictions
	close(inputs)

	// Wait for all callbacks to complete
	<-callbackDone
}
