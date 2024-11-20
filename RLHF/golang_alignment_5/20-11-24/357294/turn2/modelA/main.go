package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"
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

// Example task function that may fail
func exampleTask() error {
	// Simulate a task that fails randomly
	select {
	case <-time.After(1 * time.Second):
		return nil
	default:
		return errors.New("Task randomly failed")
	}
}

func main() {
	scheduler := NewScheduler()

	// Setup signal handling
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for sig := range c {
			log.Println("Received stop signal:", sig)
			scheduler.Cancel()
			os.Exit(0)
		}
	}()

	err := scheduler.Schedule(exampleTask, 3*time.Second)
	if err != nil {
		log.Fatal("Failed to start scheduler:", err)
	}

	// The main thread will block here waiting for signals
	log.Println("Scheduler started. Press Ctrl+C to stop...")
	<-make(chan struct{})
}
