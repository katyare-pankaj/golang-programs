package main

import (
	"fmt"
	"time"
)

type TaskScheduler struct {
	tasks    chan func() error
	errors   chan error
	quit     chan struct{}
	interval time.Duration
}

func NewTaskScheduler(interval time.Duration) *TaskScheduler {
	return &TaskScheduler{
		tasks:    make(chan func() error, 10),
		errors:   make(chan error, 10),
		quit:     make(chan struct{}),
		interval: interval,
	}
}

func (ts *TaskScheduler) Start() {
	ticker := time.NewTicker(ts.interval)
	go func() {
		for {
			select {
			case <-ticker.C:
				ts.executeTasks()
			case <-ts.quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func (ts *TaskScheduler) executeTasks() {
	for task := range ts.tasks {
		if err := task(); err != nil {
			ts.errors <- err
		}
	}
}

func (ts *TaskScheduler) AddTask(task func() error) {
	ts.tasks <- task
}

func (ts *TaskScheduler) Stop() {
	close(ts.quit)
	close(ts.tasks)
}

func (ts *TaskScheduler) ErrorHandling() {
	go func() {
		for err := range ts.errors {
			fmt.Println("Error encountered:", err)
		}
	}()
}

func main() {
	interval := 5 * time.Second
	scheduler := NewTaskScheduler(interval)

	// Sample tasks
	scheduler.AddTask(func() error {
		fmt.Println("Task 1 executed")
		return nil
	})
	scheduler.AddTask(func() error {
		fmt.Println("Task 2 encountered an error")
		return fmt.Errorf("task 2 error")
	})
	scheduler.AddTask(func() error {
		fmt.Println("Task 3 executed")
		return nil
	})

	// Start the scheduler and error handling
	scheduler.Start()
	scheduler.ErrorHandling()

	// Let the scheduler run for some time
	time.Sleep(20 * time.Second)

	// Stop the scheduler
	scheduler.Stop()
}
