package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// device represents an IoT device.
type device struct {
	name string
}

// readData simulates reading data from an IoT device.
func (d device) readData(ctx context.Context) (string, error) {
	// Sleep for a random amount of time between 1 and 5 seconds.
	delay := time.Duration(rand.Intn(5)+1) * time.Second
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(delay):
		return fmt.Sprintf("%s: %d", d.name, rand.Intn(100)), nil
	}
}

// processData simulates processing the read data from the IoT device.
func (d device) processData(data string) string {
	// Simulate processing the data.
	return fmt.Sprintf("Processed: %s", data)
}

// work starts reading data from the device, processing it, and sending it to the results channel.
func work(ctx context.Context, wg *sync.WaitGroup, device device, results chan<- string) {
	defer wg.Done()
	for {
		data, err := device.readData(ctx)
		if err != nil {
			if err == context.Canceled {
				fmt.Printf("Task for %s canceled.\n", device.name)
			} else {
				fmt.Printf("Error reading data from %s: %v\n", device.name, err)
			}
			return
		}
		processedData := device.processData(data)
		select {
		case results <- processedData:
		case <-ctx.Done():
			fmt.Printf("Task for %s canceled while sending result.\n", device.name)
			return
		}
	}
}

func main() {
	// Create a context that can be canceled.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a channel to receive interrupt signals.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c // Block until a signal is received.
		fmt.Println("User interrupted, canceling tasks...")
		cancel() // Cancel the context when a signal is received.
	}()

	var wg sync.WaitGroup
	results := make(chan string, 100) // Buffered channel to hold results.

	// Start multiple independent tasks for different devices.
	devices := []device{
		{name: "Device1"},
		{name: "Device2"},
		{name: "Device3"},
		{name: "Device4"},
	}

	for _, device := range devices {
		wg.Add(1)
		go work(ctx, &wg, device, results)
	}

	// Process the results in the main goroutine until all tasks are complete.
	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println("Reading and processing data from devices...")
	for result := range results {
		fmt.Println(result)
	}

	fmt.Println("All tasks completed.")
}
