package main

import (
	"fmt"
	"sync"
	"time"
)

// Task represents a task to be scheduled
type Task func() error

// TaskScheduler manages scheduled tasks
type TaskScheduler struct {
	jobs    map[string]Task // Map of job names to task functions
	jobCh   chan string     // Channel to add new jobs
	doneCh  chan struct{}   // Channel to signal shutdown
	mu      sync.Mutex      // Mutex for locking the jobs map
	stopCh  chan struct{}   // Channel to signal stopping the ticker
	stopped bool            // Flag to indicate if scheduler is stopped
}

// NewTaskScheduler initializes a new task scheduler
func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{
		jobs:    make(map[string]Task),
		jobCh:   make(chan string),
		doneCh:  make(chan struct{}),
		stopCh:  make(chan struct{}),
		stopped: false,
	}
}

// Schedule adds a new job to the scheduler
func (s *TaskScheduler) Schedule(name string, interval time.Duration, task Task) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.jobs[name] = task
	go s.startTicker(name, interval)
}

// startTicker starts a timer for a specific job
func (s *TaskScheduler) startTicker(name string, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			s.runTask(name)
		case <-s.stopCh:
			return
		}
	}
}

// runTask runs the scheduled task and handles errors
func (s *TaskScheduler) runTask(name string) {
	if task, ok := s.jobs[name]; ok {
		err := task()
		if err != nil {
			fmt.Printf("Error running task %s: %v\n", name, err)
		}
	} else {
		fmt.Printf("Task %s not found\n", name)
	}
}

// Stop stops the scheduler and removes all tasks
func (s *TaskScheduler) Stop() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.stopped {
		s.stopped = true
		close(s.stopCh)
		close(s.jobCh)
		<-s.doneCh
		for name := range s.jobs {
			delete(s.jobs, name)
		}
	}
}

// Run starts the scheduler
func (s *TaskScheduler) Run() {
	for {
		select {
		case name := <-s.jobCh:
			s.Schedule(name, time.Minute, func() error {
				return fmt.Errorf("Running periodic task: %s", name)
			})
		case <-s.stopCh:
			close(s.doneCh)
			return
		}
	}
}

func main() {
	scheduler := NewTaskScheduler()

	go scheduler.Run()

	scheduler.Schedule("task1", time.Second*5, func() error {
		fmt.Println("Running task1 every 5 seconds")
		return nil
	})

	scheduler.Schedule("task2", time.Second*10, func() error {
		fmt.Println("Running task2 every 10 seconds")
		return nil
	})

	time.Sleep(time.Second * 20) // Let the tasks run for a while

	scheduler.Stop()
}
