package main

import (
	"fmt"
	"sync"
	"time"
)

// Callback function type for data processing
type Callback func(data int) int

// DataProcessor structure that holds callbacks and sync structures
type DataProcessor struct {
	callbacks []Callback
	mutex     sync.Mutex
	wg        sync.WaitGroup
}

// NewDataProcessor initializes and returns a new DataProcessor
func NewDataProcessor() *DataProcessor {
	return &DataProcessor{
		callbacks: []Callback{},
	}
}

// RegisterCallback adds a new callback function to the list of callbacks
func (dp *DataProcessor) RegisterCallback(callback Callback) {
	dp.mutex.Lock()
	defer dp.mutex.Unlock()
	dp.callbacks = append(dp.callbacks, callback)
}

// ProcessData runs the callbacks concurrently on incoming data
func (dp *DataProcessor) ProcessData(dataStream <-chan int, results chan<- int) {
	for data := range dataStream {
		dp.wg.Add(1)
		go dp.process(data, results)
	}
	dp.wg.Wait()
	close(results)
}

// process is a helper function to apply callbacks to data
func (dp *DataProcessor) process(data int, results chan<- int) {
	defer dp.wg.Done()
	dp.mutex.Lock()
	defer dp.mutex.Unlock()
	for _, callback := range dp.callbacks {
		data = callback(data)
	}
	results <- data
}

// Example callback function that increments the data by 1
func incrementCallback(data int) int {
	return data + 1
}

// Example callback function that doubles the data
func doubleCallback(data int) int {
	return data * 2
}

func main() {
	// Initialize the data processor
	dataProcessor := NewDataProcessor()

	// Register some callback functions
	dataProcessor.RegisterCallback(incrementCallback)
	dataProcessor.RegisterCallback(doubleCallback)

	// Create channels for data stream and results
	dataStream := make(chan int)
	results := make(chan int)

	// Run the data processor
	go dataProcessor.ProcessData(dataStream, results)

	// Simulate real-time data input
	go func() {
		for i := 0; i < 10; i++ {
			dataStream <- i
			time.Sleep(100 * time.Millisecond) // simulate delay in data arrival
		}
		close(dataStream)
	}()

	// Collect and print results
	for result := range results {
		fmt.Println(result)
	}
}
