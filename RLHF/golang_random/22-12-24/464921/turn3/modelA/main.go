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

// Task represents a task that can be executed concurrently.
type Task struct {
	name    string
	delay   time.Duration
	cancel  context.CancelFunc
	results chan string
}

// NewTask creates a new task with a random delay.
func NewTask(name string) *Task {
	rand.Seed(time.Now().UnixNano())
	delay := time.Duration(rand.Intn(5)+1) * time.Second
	ctx, cancel := context.WithCancel(context.Background())
	return &Task{
		name:    name,
		delay:   delay,
		cancel:  cancel,
		results: make(chan string, 100),
	}
}

// Work performs the task operations.
func (t *Task) Work(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Task %s canceled.\n", t.name)
			return
		default:
			time.Sleep(t.delay)
			fmt.Printf("Task %s: operation complete.\n", t.name)
			select {
			case <-ctx.Done():
				fmt.Printf("Task %s canceled while processing results.\n", t.name)
				return
			case results := <-t.results:
				fmt.Printf("Task %s: %s\n", t.name, results)
			}
		}
	}
}

// SendResult sends a result to the task's results channel.
func (t *Task) SendResult(result string) {
	t.results <- result
}

func main() {
	tasks := []*Task{
		NewTask("Task1"),
		NewTask("Task2"),
		NewTask("Task3"),
	}

	var wg sync.WaitGroup
	cancelCtx, cancel := context.WithCancel(context.Background())

	// Signal handling to cancel tasks upon interruption.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c // Wait for a signal.
		fmt.Println("User interrupted, canceling tasks...")
		cancel() // Cancel the context to stop the tasks.
	}()

	for _, task := range tasks {
		wg.Add(1)
		go task.Work(cancelCtx, &wg)
		go func(task *Task) {
			time.Sleep(5 * time.Second)
			task.SendResult(fmt.Sprintf("Result from %s", task.name))
		}(task)
	}

	// Block the main goroutine until all tasks are finished.
	fmt.Println("Starting tasks...")
	go func() {
		wg.Wait()
		close(c)
	}()

	select {
	case <-time.After(10 * time.Second):
		fmt.Println("All tasks completed successfully.")
	}
}
