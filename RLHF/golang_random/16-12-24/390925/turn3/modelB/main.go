package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

// Task interface defines the behaviour all tasks must follow
type Task interface {
	Run()
}

// PrintTask is a simple task that prints a message
type PrintTask struct {
	Message string
}

func (t *PrintTask) Run() {
	fmt.Println("Running PrintTask:", t.Message)
	time.Sleep(1 * time.Second)
}

// CalculateTask is a task that performs a calculation
type CalculateTask struct {
	A, B int
}

func (t *CalculateTask) Run() {
	result := t.A + t.B
	fmt.Println("Running CalculateTask:", t.A, "+", t.B, "=", result)
	time.Sleep(1 * time.Second)
}

// LoadAndExecuteTasks uses reflection to dynamically load and execute tasks
func LoadAndExecuteTasks(taskDescs []string, taskExecutor chan Task) {
	for _, taskDesc := range taskDescs {
		parts := reflect.Split(taskDesc, ".")
		if len(parts) != 2 {
			fmt.Println("Invalid task description:", taskDesc)
			continue
		}

		pkg := parts[0]
		typ := parts[1]

		// Import the package dynamically
		if err := importPackage(pkg); err != nil {
			fmt.Println("Error importing package", pkg, ":", err)
			continue
		}

		// Create an instance of the task using reflection
		var task Task
		switch typ {
		case "PrintTask":
			task = &PrintTask{Message: "Dynamic Task"}
		case "CalculateTask":
			task = &CalculateTask{A: 42, B: 13}
		default:
			fmt.Println("Unknown task type:", typ)
			continue
		}

		// Send the task to the executor
		taskExecutor <- task
	}
	close(taskExecutor)
}

func importPackage(pkg string) error {
	// For demonstration purposes, we're not actually importing dynamically
	// In a real scenario, you'd use a package like "plugin" to dynamically load packages.
	return nil
}

// TaskExecutor handles executing tasks concurrently
func TaskExecutor(taskExecutor chan Task) {
	var wg sync.WaitGroup

	// Create a fixed number of worker goroutines
	const numWorkers = 5
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range taskExecutor {
				task.Run()
			}
		}()
	}

	wg.Wait()
}

func main() {
	// Define task descriptions (package.Type)
	taskDescs := []string{
		"main.PrintTask",
		"main.CalculateTask",
		"main.PrintTask",
		"main.CalculateTask",
	}

	// Create a channel for tasks
	taskExecutor := make(chan Task)

	// Start task loading and execution concurrently
	go LoadAndExecuteTasks(taskDescs, taskExecutor)
	go TaskExecutor(taskExecutor)
}
