package main

import (
	"fmt"
	"time"
)

type Task struct {
	ID     int
	Data   string
	Source int
}
type RateLimiter struct {
	burst  int
	period time.Duration
	ch     chan Task
}

func NewRateLimiter(burst int, period time.Duration) *RateLimiter {
	return &RateLimiter{
		burst:  burst,
		period: period,
		ch:     make(chan Task, burst),
	}
}
func (rl *RateLimiter) ProcessTasks() {
	ticker := time.NewTicker(rl.period)
	defer ticker.Stop()
	for {
		select {
		case task := <-rl.ch:
			fmt.Printf("Processing Task %d: %s (Source: %d)\n", task.ID, task.Data, task.Source)
			time.Sleep(time.Second) // Simulate task processing time
		case <-ticker.C:
			for len(rl.ch) > 0 {
				task := <-rl.ch
				fmt.Printf("Processing Task %d: %s (Source: %d)\n", task.ID, task.Data, task.Source)
				time.Sleep(time.Second) // Simulate task processing time
			}
		}
	}
}
func GenerateTasks(source int, out chan<- Task, rateLimiter *RateLimiter) {
	taskID := 1
	for {
		task := Task{ID: taskID, Data: fmt.Sprintf("Task from source %d", source), Source: source}
		rateLimiter.ch <- task
		taskID++
	}
}

func main() {
	// Create a rate limiter with a burst of 2 tasks and a period of 2 seconds
	rateLimiter := NewRateLimiter(2, 2*time.Second)

	// Start the rate limiter
	go rateLimiter.ProcessTasks()

	// Generate tasks from different sources (channels)
	taskChannel1 := make(chan Task)
	go GenerateTasks(1, taskChannel1, rateLimiter)
	taskChannel2 := make(chan Task)
	go GenerateTasks(2, taskChannel2, rateLimiter)

	// Simulate processing tasks for 10 seconds
	time.Sleep(10 * time.Second)
	fmt.Println("Task processing completed.")
}
