package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Task represents a job to be executed at intervals
type Task struct {
	ID      int
	Message string
}

// jobScheduler handles scheduling tasks at a specified interval
func jobScheduler(ctx context.Context, wg *sync.WaitGroup, interval time.Duration, taskChannel <-chan Task) {
	defer wg.Done()

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Shutting down job scheduler...")
			return
		case <-ticker.C:
			select {
			case task := <-taskChannel:
				fmt.Printf("Executing Task %d: %s\n", task.ID, task.Message)
			default:
				fmt.Println("No tasks to execute at this interval.")
			}
		}
	}
}

func main() {
	taskChannel := make(chan Task)
	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure context is canceled when main exits

	// Start the job scheduler
	wg.Add(1)
	go jobScheduler(ctx, &wg, 2*time.Second, taskChannel)

	// Send tasks to the task channel
	go func() {
		for i := 1; i <= 5; i++ {
			taskChannel <- Task{ID: i, Message: fmt.Sprintf("Task-%d payload", i)}
			time.Sleep(1 * time.Second) // Delay between sending tasks
		}

		// Close the task channel to indicate no more tasks will be sent
		close(taskChannel)
	}()

	// Wait a while before shutting down to allow some tasks to be processed
	time.Sleep(10 * time.Second)

	// Initiate graceful shutdown
	cancel()  // Signal the scheduler to stop
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All tasks processed and scheduler shut down gracefully.")
}
