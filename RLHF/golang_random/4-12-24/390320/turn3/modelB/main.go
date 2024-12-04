package main

import (
	"fmt"
	"time"
)

// Task represents a unit of work to be executed.
type Task struct {
	ID   int
	Exec func() error // Exec is the function that performs the task.
}

// TaskScheduler schedules and executes tasks.
type TaskScheduler struct {
	tasks      []Task
	errorChan  chan error
	execPeriod time.Duration
	stopChan   chan struct{}
}

// NewTaskScheduler initializes a TaskScheduler with a given execution period.
func NewTaskScheduler(execPeriod time.Duration) *TaskScheduler {
	return &TaskScheduler{
		tasks:      []Task{},
		errorChan:  make(chan error, 10),
		execPeriod: execPeriod,
		stopChan:   make(chan struct{}),
	}
}

// AddTask adds a new task to the scheduler.
func (ts *TaskScheduler) AddTask(task Task) {
	ts.tasks = append(ts.tasks, task)
}

// Start begins the task execution at regular intervals.
func (ts *TaskScheduler) Start() {
	ticker := time.NewTicker(ts.execPeriod)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				ts.runTasks()
			case <-ts.stopChan:
				close(ts.stopChan) // Ensure the stop channel is closed to prevent leaks.
				return
			}
		}
	}()
}

// runTasks executes each task concurrently, capturing any errors.
func (ts *TaskScheduler) runTasks() {
	var doneCount int
	done := make(chan struct{})
	for _, task := range ts.tasks {
		go func(t Task) {
			if err := t.Exec(); err != nil {
				ts.errorChan <- fmt.Errorf("task %d: %w", t.ID, err)
			}
			done <- struct{}{}
		}(task)
	}

	for range ts.tasks {
		<-done
		doneCount++
		if doneCount == len(ts.tasks) {
			close(done)
		}
	}
}

// Stop signals the scheduler to stop executing tasks.
func (ts *TaskScheduler) Stop() {
	ts.stopChan <- struct{}{}
}

// ErrorHandling logs the errors captured during task execution.
func (ts *TaskScheduler) ErrorHandling() {
	go func() {
		for err := range ts.errorChan {
			fmt.Println("Error encountered:", err)
		}
	}()
}

func main() {
	execPeriod := 5 * time.Second // Define execution period interval.
	scheduler := NewTaskScheduler(execPeriod)

	// Define example tasks with unique IDs and execution functions.
	scheduler.AddTask(Task{ID: 1, Exec: func() error {
		fmt.Println("Executing Task 1")
		return nil
	}})
	scheduler.AddTask(Task{ID: 2, Exec: func() error {
		fmt.Println("Executing Task 2")
		return fmt.Errorf("something went wrong in Task 2")
	}})
	scheduler.AddTask(Task{ID: 3, Exec: func() error {
		fmt.Println("Executing Task 3")
		return nil
	}})

	// Start the scheduler and handle errors.
	scheduler.Start()
	scheduler.ErrorHandling()

	// Let the scheduler run for some time.
	time.Sleep(20 * time.Second)

	// Stop the scheduler.
	scheduler.Stop()

	// Give time to handle cleanup.
	time.Sleep(2 * time.Second)
}
