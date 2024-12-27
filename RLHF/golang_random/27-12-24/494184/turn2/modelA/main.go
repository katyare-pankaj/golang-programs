package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type task struct {
	id       int
	data     string
	duration time.Duration
}

func (t *task) execute() (string, error) {
	fmt.Printf("Worker processing task ID: %d\n", t.id)
	time.Sleep(t.duration) // Simulate work
	return fmt.Sprintf("Result for task %d: processed %s", t.id, t.data), nil
}

func worker(wg *sync.WaitGroup, taskChannel <-chan *task, resultChannel chan<- *task) {
	defer wg.Done()
	for {
		select {
		case task, ok := <-taskChannel:
			if !ok {
				// Channel is closed, stop worker
				return
			}
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second) // Set 2-second timeout
			defer cancel()
			select {
			case <-ctx.Done(): // Timeout occurred
				resultChannel <- &task{id: task.id, data: task.data, duration: task.duration}
				fmt.Printf("Timeout exceeded for task ID: %d\n", task.id)
			default:
				result, err := task.execute()
				if err != nil {
					resultChannel <- &task{id: task.id, data: "", duration: task.duration, err: err}
				} else {
					resultChannel <- &task{id: task.id, data: result, duration: task.duration}
				}
			}
		}
	}
}

func main() {
	numWorkers := 3
	taskChannel := make(chan *task)
	resultChannel := make(chan *task)
	var wg sync.WaitGroup

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(&wg, taskChannel, resultChannel)
	}

	tasks := []*task{
		&task{id: 1, data: "hello", duration: 1 * time.Second},
		&task{id: 2, data: "world", duration: 4 * time.Second},
		&task{id: 3, data: "work", duration: 2 * time.Second},
	}

	for _, t := range tasks {
		taskChannel <- t
	}

	close(taskChannel) // Signal workers to stop

	for i := range tasks {
		result := <-resultChannel
		if result.err != nil {
			fmt.Printf("Error processing task ID: %d - %v\n", result.id, result.err)
		} else {
			fmt.Printf("Completed task ID: %d with result: %s\n", result.id, result.data)
		}
	}

	wg.Wait()
}
