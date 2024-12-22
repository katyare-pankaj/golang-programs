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

// Device represents an IoT device.
type Device struct {
	Name string
}

// ReadData simulates reading data from an IoT device.
func (d Device) ReadData(ctx context.Context) (string, error) {
	delay := time.Duration(rand.Intn(5)+1) * time.Second
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(delay):
		return fmt.Sprintf("%s: %d", d.Name, rand.Intn(100)), nil
	}
}

// ProcessData simulates processing the read data from the IoT device.
func (d Device) ProcessData(data string) string {
	// Simulate processing the data.
	time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
	return fmt.Sprintf("Processed: %s", data)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\nUser interrupted, canceling tasks...")
		cancel()
	}()

	var wg sync.WaitGroup
	results := make(chan string, 100)

	devices := []Device{
		{Name: "Device1"},
		{Name: "Device2"},
		{Name: "Device3"},
		{Name: "Device4"},
	}

	for _, device := range devices {
		wg.Add(1)
		go func(d Device) {
			defer wg.Done()
			for {
				data, err := d.ReadData(ctx)
				if err != nil {
					if err == context.Canceled {
						fmt.Printf("Task for %s canceled.\n", d.Name)
					} else {
						fmt.Printf("Error reading data from %s: %v\n", d.Name, err)
					}
					return
				}
				processedData := d.ProcessData(data)
				select {
				case results <- processedData:
				case <-ctx.Done():
					fmt.Printf("Task for %s canceled while sending result.\n", d.Name)
					return
				}
			}
		}(device)
	}

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
