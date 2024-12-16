package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

// Task represents a generic task to be executed
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

// TaskExecutor uses reflection to execute tasks
func TaskExecutor(tasks chan Task) {
	for task := range tasks {
		val := reflect.ValueOf(task)
		typ := val.Type()

		fmt.Printf("Executing task of type %s\n", typ.Name())

		// Call the Run method using reflection
		val.MethodByName("Run").Call(nil)
	}
}

func main() {
	// Create a channel for tasks
	tasks := make(chan Task)

	// Start a goroutine to execute tasks
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		TaskExecutor(tasks)
		wg.Done()
	}()

	// Dispatch tasks concurrently
	tasks <- &PrintTask{Message: "Hello, World!"}
	tasks <- &CalculateTask{A: 10, B: 20}
	tasks <- &PrintTask{Message: "Task completed."}
	tasks <- &CalculateTask{A: 100, B: 200}

	// Close the task channel to signal the executor to stop
	close(tasks)

	wg.Wait()
}
