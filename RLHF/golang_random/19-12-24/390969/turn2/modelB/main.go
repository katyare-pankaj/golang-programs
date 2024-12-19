package main

import (
	"container/heap"
	"fmt"
	"sync"
	"time"
)

// PriorityTask represents a task with an associated priority
type PriorityTask struct {
	Priority int
	Task     Task
}

// Task represents a task to be processed
type Task struct {
	ID   int
	Data string
}

// PriorityTaskQueue implements heap.Interface to prioritize tasks
type PriorityTaskQueue []PriorityTask

func (pq PriorityTaskQueue) Len() int { return len(pq) }

func (pq PriorityTaskQueue) Less(i, j int) bool {
	// Return true if the priority of the first task is greater than the priority of the second task
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityTaskQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityTaskQueue) Push(x interface{}) {
	task := x.(PriorityTask)
	*pq = append(*pq, task)
}

func (pq *PriorityTaskQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	task := old[n-1]
	*pq = old[0 : n-1]
	return task
}

// Worker processes tasks from the work channel
func Worker(wg *sync.WaitGroup, work chan PriorityTask) {
	defer wg.Done()
	for task := range work {
		fmt.Printf("Processing task %d with priority %d\n", task.Task.ID, task.Priority)
		// Simulate task processing time
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// Create a priority task queue
	pq := &PriorityTaskQueue{}
	heap.Init(pq)

	var wg sync.WaitGroup
	numWorkers := 3 // Adjust the number of workers as needed

	// Start worker goroutines
	work := make(chan PriorityTask, 100) // Buffer the work channel
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go Worker(&wg, work)
	}

	// Assign tasks with different priorities to the work channel
	tasks := []PriorityTask{
		{Priority: 1, Task: Task{ID: 1, Data: "High priority task"}},
		{Priority: 5, Task: Task{ID: 2, Data: "Medium priority task"}},
		{Priority: 3, Task: Task{ID: 3, Data: "Low priority task"}},
		{Priority: 2, Task: Task{ID: 4, Data: "Another medium priority task"}},
		{Priority: 1, Task: Task{ID: 5, Data: "Yet another high priority task"}},
	}

	for _, task := range tasks {
		heap.Push(pq, task)
	}

	// Dispatch tasks to workers in order of priority
	for pq.Len() > 0 {
		task := heap.Pop(pq).(PriorityTask)
		work <- task
	}

	// Close the work channel to signal workers to exit
	close(work)
	wg.Wait()
	fmt.Println("All tasks completed.")
}
