package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Device represents an IoT device
type Device struct {
	Name string
}

// ReadData simulates reading data from an IoT device.
func (d Device) ReadData(ctx context.Context) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err() // Return context error if it is canceled or timed out
	default:
		// Simulate a random delay to read data
		delay := time.Duration(rand.Intn(500)) * time.Millisecond
		time.Sleep(delay)
		return fmt.Sprintf("%s: %d", d.Name, rand.Intn(100)), nil
	}
}

// ProcessData simulates processing the read data from the IoT device.
func (d Device) ProcessData(data string) string {
	// Simulate processing the data
	return fmt.Sprintf("Processed: %s", data)
}

func main() {
	// Create a context with a timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a slice of IoT devices
	devices := []Device{
		{Name: "Device1"},
		{Name: "Device2"},
		{Name: "Device3"},
		{Name: "Device4"},
	}

	var wg sync.WaitGroup
	results := make(chan string, len(devices))

	// Start a goroutine for each device to read data and process it
	for _, device := range devices {
		wg.Add(1)
		go func(d Device) {
			defer wg.Done()
			data, err := d.ReadData(ctx)
			if err != nil {
				fmt.Printf("Error reading data from %s: %v\n", d.Name, err)
				return
			}
			processedData := d.ProcessData(data)
			results <- processedData
		}(device)
	}

	// Wait for all tasks to complete
	go func() {
		wg.Wait()
		close(results)
	}()

	// Print the processed results
	for result := range results {
		fmt.Println(result)
	}

	// Check if the context timed out
	if err := ctx.Err(); err == context.DeadlineExceeded {
		fmt.Println("Context timed out. Not all devices may have been processed.")
	}
}
