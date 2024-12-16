package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

// Task interface represents a generic task that can be executed
type Task interface {
	Run()
}

// PrintTask is a task that prints a message
type PrintTask struct {
	Message string
}

func (t *PrintTask) Run() {
	fmt.Println("Executing PrintTask:", t.Message)
	time.Sleep(1 * time.Second) // Simulating work
}

// CalculateTask is a task that performs a calculation
type CalculateTask struct {
	A, B int
}

func (t *CalculateTask) Run() {
	result := t.A + t.B
	fmt.Printf("Executing CalculateTask: %d + %d = %d\n", t.A, t.B, result)
	time.Sleep(1 * time.Second) // Simulating work
}

// TaskExecutor is responsible for executing tasks concurrently
func TaskExecutor(tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done() // Notify that this goroutine is done when returning

	for task := range tasks {
		// Use reflection to call Run method
		val := reflect.ValueOf(task)
		val.MethodByName("Run").Call(nil)
	}
}

func main() {
	// Create a channel for tasks
	tasks := make(chan Task)

	// WaitGroup to manage goroutines
	var wg sync.WaitGroup
	wg.Add(1)

	// Start a goroutine to execute tasks
	go TaskExecutor(tasks, &wg)

	// Dispatch tasks to the channel
	tasks <- &PrintTask{Message: "Hello, World!"}
	tasks <- &CalculateTask{A: 10, B: 20}
	tasks <- &PrintTask{Message: "Executing another task!"}
	tasks <- &CalculateTask{A: 100, B: 200}

	// Close the task channel to indicate no more tasks will be sent
	close(tasks)

	// Wait for the task executor to finish
	wg.Wait()
}
