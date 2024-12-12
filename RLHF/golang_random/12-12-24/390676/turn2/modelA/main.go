package main

import (
	"context"
	"log"
	"sync"
	"time"
)

// Task represents a unit of work to be scheduled
type Task struct {
	ID          int
	Interval    time.Duration
	ExecuteFunc func() error
}

// Scheduler manages task scheduling
type Scheduler struct {
	tasks  []Task
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

// NewScheduler initializes a new Scheduler
func NewScheduler() *Scheduler {
	ctx, cancel := context.WithCancel(context.Background())
	return &Scheduler{
		tasks:  []Task{},
		ctx:    ctx,
		cancel: cancel,
	}
}

// AddTask adds a new task to the scheduler
func (s *Scheduler) AddTask(task Task) {
	s.tasks = append(s.tasks, task)
}

// Start begins executing all scheduled tasks
func (s *Scheduler) Start() {
	for _, task := range s.tasks {
		s.wg.Add(1)
		go s.runTask(task)
	}
}

// runTask executes the provided task at specified intervals
func (s *Scheduler) runTask(task Task) {
	defer s.wg.Done()

	ticker := time.NewTicker(task.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := task.ExecuteFunc(); err != nil {
				log.Printf("Task %d failed: %v", task.ID, err)
			}
		case <-s.ctx.Done():
			log.Printf("Stopping task %d", task.ID)
			return
		}
	}
}

// Stop gracefully shuts down the scheduler and waits for all tasks to finish
func (s *Scheduler) Stop() {
	s.cancel()
	s.wg.Wait()
}

// Example task function
func exampleTaskFunc(id int) func() error {
	return func() error {
		log.Printf("Executing task %d", id)
		// Simulate work
		time.Sleep(1 * time.Second)
		return nil // or return an error if something goes wrong
	}
}

// Main application
func main() {
	scheduler := NewScheduler()

	// Adding some tasks to the scheduler
	for i := 1; i <= 3; i++ {
		scheduler.AddTask(Task{
			ID:          i,
			Interval:    2 * time.Second,
			ExecuteFunc: exampleTaskFunc(i),
		})
	}

	// Start the scheduler
	scheduler.Start()

	// Run for a while to see the periodic task execution
	time.Sleep(10 * time.Second)

	// Stop the scheduler gracefully
	scheduler.Stop()

	log.Println("Scheduler stopped")
}
