package main

import (
	"fmt"
	"sync"
)

type Event string
type Task func()
type WorkflowEngine struct {
	eventHandlers map[Event]Task
	wg            sync.WaitGroup
}

func (we *WorkflowEngine) Register(event Event, task Task) {
	we.eventHandlers[event] = task
}

func (we *WorkflowEngine) Run(event Event) {
	we.wg.Add(1)
	defer we.wg.Done()

	if task, ok := we.eventHandlers[event]; ok {
		task()
	}
}

func (we *WorkflowEngine) Wait() {
	we.wg.Wait()
}

func main() {
	// Create a new workflow engine
	we := &WorkflowEngine{
		eventHandlers: make(map[Event]Task),
	}

	// Register event handlers
	we.Register("download_completed", func() {
		fmt.Println("Download completed!")
	})

	we.Register("file_processed", func() {
		fmt.Println("File processed successfully!")
	})

	we.Register("upload_failed", func() {
		fmt.Println("Upload failed!")
	})

	// Simulate asynchronous events being triggered
	go func() {
		we.Run("download_completed")
		we.Run("file_processed")
	}()

	go func() {
		we.Run("upload_failed")
	}()

	// Wait for all background tasks to complete
	we.Wait()

	fmt.Println("All tasks completed.")
}
