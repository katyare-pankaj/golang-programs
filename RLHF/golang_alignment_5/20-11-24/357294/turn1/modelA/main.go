package main

import (
	"errors"
	"log"
	"time"
)

// Task defines a function to be scheduled
type Task func() error

// Scheduler manages the execution of tasks
type Scheduler struct {
	cancelFunc func()
}

// NewScheduler creates a new scheduler
func NewScheduler() *Scheduler {
	return &Scheduler{}
}

// Schedule sets up the task to be run every duration
func (s *Scheduler) Schedule(task Task, duration time.Duration) error {
	if s.cancelFunc != nil {
		return errors.New("scheduler already running")
	}

	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	for range ticker.C {
		err := task()
		if err != nil {
			log.Printf("Task execution failed: %v\n", err)
		} else {
			log.Println("Task executed successfully.")
		}
	}

	return nil
}

// Cancel stops the scheduling of tasks
func (s *Scheduler) Cancel() {
	if s.cancelFunc != nil {
		s.cancelFunc()
	}
}

func main() {
	// Example task function that may fail
	task := func() error {
		// Simulate task failure 50% of the time
		select {
		case <-time.After(1 * time.Second):
			return nil
		default:
			return errors.New("Task randomly failed")
		}
	}

	scheduler := NewScheduler()
	err := scheduler.Schedule(task, 5*time.Second)
	if err != nil {
		log.Fatal("Failed to start scheduler:", err)
	}

	// Simulate workload completion, cancel the scheduler
	time.Sleep(15 * time.Second)
	scheduler.Cancel()
}
