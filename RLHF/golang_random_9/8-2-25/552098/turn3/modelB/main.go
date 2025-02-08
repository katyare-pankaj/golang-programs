package main

import (
	"fmt"
	"sync"
	"time"
)

// Job represents a function to be scheduled.
type Job func()

// Scheduler runs jobs at specified intervals.
type Scheduler struct {
	mu   sync.Mutex
	jobs map[string]*jobEntry
	quit chan struct{}
}

type jobEntry struct {
	job      Job
	interval time.Duration
	ticker   *time.Ticker
}

// NewScheduler creates a new Scheduler.
func NewScheduler() *Scheduler {
	return &Scheduler{
		jobs: make(map[string]*jobEntry),
		quit: make(chan struct{}),
	}
}

// Start begins running the scheduler.
func (s *Scheduler) Start() {
	go s.run()
}

// Stop gracefully shuts down the scheduler, waiting for ongoing jobs to complete.
func (s *Scheduler) Stop() {
	close(s.quit)
}

// AddJob adds a job to the scheduler with a given interval.
func (s *Scheduler) AddJob(name string, job Job, interval time.Duration) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.jobs[name]; ok {
		return fmt.Errorf("job '%s' already exists", name)
	}

	entry := &jobEntry{
		job:      job,
		interval: interval,
		ticker:   time.NewTicker(interval),
	}
	s.jobs[name] = entry

	go func() {
		for {
			select {
			case <-entry.ticker.C:
				entry.job()
			case <-s.quit:
				entry.ticker.Stop()
				return
			}
		}
	}()

	return nil
}

// RemoveJob removes a job from the scheduler.
func (s *Scheduler) RemoveJob(name string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	entry, ok := s.jobs[name]
	if !ok {
		return fmt.Errorf("job '%s' not found", name)
	}

	delete(s.jobs, name)
	entry.ticker.Stop()

	return nil
}

func (s *Scheduler) run() {
	for {
		select {
		case <-s.quit:
			s.mu.Lock()
			for _, entry := range s.jobs {
				entry.ticker.Stop()
			}
			s.mu.Unlock()
			return
		}
	}
}

func job1() {
	fmt.Println("Running job 1...")
	time.Sleep(time.Second * 2)
	fmt.Println("Job 1 completed.")
}

func job2() {
	fmt.Println("Running job 2...")
	time.Sleep(time.Second * 3)
	fmt.Println("Job 2 completed.")
}

func main() {
	scheduler := NewScheduler()
	scheduler.Start()

	// Add jobs
	scheduler.AddJob("job1", job1, time.Second*3)
	scheduler.AddJob("job2", job2, time.Second*5)

	time.Sleep(time.Second * 10) // Run the jobs for 10 seconds

	// Remove a job
	scheduler.RemoveJob("job2")

	time.Sleep(time.Second * 5) // Give some time for the remaining jobs to finish

	// Gracefully shutdown
	scheduler.Stop()
	fmt.Println("Scheduler shut down.")
}
