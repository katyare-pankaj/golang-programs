package main

import (
	"fmt"
	"sync"
	"time"
)

// Task represents a task to be executed
type Task func() error

// TaskScheduler manages a list of tasks to be executed periodically
type TaskScheduler struct {
	tasks    []Task
	errors   chan error
	quit     chan struct{}
	interval time.Duration
	wg       sync.WaitGroup
}

// NewTaskScheduler initializes a new TaskScheduler
func NewTaskScheduler(interval time.Duration) *TaskScheduler {
	return &TaskScheduler{
		tasks:    make([]Task, 0),
		errors:   make(chan error, 10),
		quit:     make(chan struct{}),
		interval: interval,
	}
}

// AddTask adds a new task to the scheduler
func (ts *TaskScheduler) AddTask(task Task) {
	ts.tasks = append(ts.tasks, task)
}

// Start begins executing the tasks periodically
func (ts *TaskScheduler) Start() {
	ticker := time.NewTicker(ts.interval)
	go func() {
		defer ticker.Stop()
		defer ts.wg.Done()
		for {
			select {
			case <-ticker.C:
				ts.executeTasks()
			case <-ts.quit:
				return
			}
		}
	}()
}

// Stop stops the scheduler
func (ts *TaskScheduler) Stop() {
	close(ts.quit)
	ts.wg.Wait()
}

// executeTasks runs all the tasks concurrently
func (ts *TaskScheduler) executeTasks() {
	for _, task := range ts.tasks {
		ts.wg.Add(1)
		go func(task Task) {
			defer ts.wg.Done()
			if err := task(); err != nil {
				ts.errors <- fmt.Errorf("error in task: %w", err)
			}
		}(task)
	}
	ts.wg.Wait()
}

// collectErrors collects errors that occurred during task execution
func (ts *TaskScheduler) collectErrors() {
	for err := range ts.errors {
		fmt.Println("Task execution error:", err)
	}
}

func main() {
	scheduler := NewTaskScheduler(5 * time.Second)

	// Add some sample tasks to the scheduler
	scheduler.AddTask(func() error {
		fmt.Println("Task 1 executing...")
		// Simulate a potential error
		return fmt.Errorf("simulated error in Task 1")
	})
	scheduler.AddTask(func() error {
		fmt.Println("Task 2 executing...")
		time.Sleep(2 * time.Second) // Simulate a time-consuming task
		return nil
	})
	scheduler.AddTask(func() error {
		fmt.Println("Task 3 executing...")
		return nil
	})

	// Start the scheduler and error handling
	go scheduler.Start()
	go scheduler.collectErrors()

	// Let the program run for some time
	time.Sleep(30 * time.Second)

	// Stop the scheduler
	scheduler.Stop()
	fmt.Println("Main program ended.")
}
