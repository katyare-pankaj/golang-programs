package main

import (
	"fmt"
	"os"
)

func main() {
	if err := performTask(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("Task completed successfully")
}

func performTask() error {
	// Simulating opening a resource
	resource, err := openResource("example")
	if err != nil {
		return fmt.Errorf("failed to open resource: %w", err)
	}

	// Ensure that the resource will be closed when the function exits
	defer closeResource(resource)

	// Simulate some task that might fail
	if err := doSomeWork(resource); err != nil {
		return fmt.Errorf("error during work: %w", err)
	}

	// Simulate successful completion
	return nil
}

func openResource(name string) (*os.File, error) {
	fmt.Printf("Opening resource: %s\n", name)
	// Simulating opening a file as a resource for demonstration
	file, err := os.Create(name + ".txt")
	if err != nil {
		return nil, err
	}
	return file, nil
}

func closeResource(file *os.File) {
	fmt.Printf("Closing resource: %s\n", file.Name())
	file.Close()
}

func doSomeWork(file *os.File) error {
	// Simulate a failure
	return fmt.Errorf("failed operation")
}
