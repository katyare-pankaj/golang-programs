package main

import (
	"container/heap"
	"errors"
	"fmt"
)

// PriorityTask represents a task with an associated priority.
type PriorityTask struct {
	Task     interface{}
	Priority int
}

// BoundedQueue represents a queue with dynamic resizing and priority-based enqueueing.
type BoundedQueue struct {
	tasks []PriorityTask
	limit int
}

// Implement heap interface methods
func (bq *BoundedQueue) Len() int           { return len(bq.tasks) }
func (bq *BoundedQueue) Less(i, j int) bool { return bq.tasks[i].Priority < bq.tasks[j].Priority }
func (bq *BoundedQueue) Swap(i, j int)      { bq.tasks[i], bq.tasks[j] = bq.tasks[j], bq.tasks[i] }
func (bq *BoundedQueue) Push(x interface{}) {
	task := x.(PriorityTask)
	bq.tasks = append(bq.tasks, task)
	heap.Fix(bq, bq.Len()-1) // Restore heap property after pushing
}
func (bq *BoundedQueue) Pop() interface{} {
	old := bq.tasks
	n := len(old)
	task := old[0]
	bq.tasks = old[1:n]
	heap.Fix(bq, 0) // Restore heap property after popping
	return task
}

// NewBoundedQueue creates a new BoundedQueue with the specified initial limit.
func NewBoundedQueue(initialLimit int) *BoundedQueue {
	return &BoundedQueue{
		tasks: make([]PriorityTask, 0),
		limit: initialLimit,
	}
}

// Enqueue adds an element to the queue with a given priority.
// The queue will automatically resize if the limit is reached.
func (bq *BoundedQueue) Enqueue(task interface{}, priority int) {
	heap.Push(bq, PriorityTask{Task: task, Priority: priority})
	if bq.Len() > bq.limit {
		bq.limit = bq.Len()
	}
}

// Dequeue removes and returns the element with the highest priority from the queue.
// Returns an error if the queue is empty.
func (bq *BoundedQueue) Dequeue() (interface{}, error) {
	if len(bq.tasks) == 0 {
		return nil, errors.New("queue is empty")
	}
	return heap.Pop(bq).(PriorityTask).Task, nil
}

// Resize adjusts the limit of the BoundedQueue.
func (bq *BoundedQueue) Resize(newLimit int) {
	bq.limit = newLimit
}

func main() {
	bq := NewBoundedQueue(3)
	heap.Init(bq)
	
	// Enqueue tasks with priorities
	bq.Enqueue("Task 1", 3)
	bq.Enqueue("Task 3", 1)
	bq.Enqueue("Task 2", 2)
	
	fmt.Println("Task Priority:")
	for bq.Len() > 0 {
		task, _ := bq.Dequeue()
		fmt.Println(task)
	}
	
	// Resize the queue and enqueue more tasks
	bq.Resize(5)
	bq.Enqueue("Task 4", 4)
	bq.Enqueue("Task 5", 5)
	bq.Enqueue("Task 6", 6)
	
	fmt.Println("\nTask Priority after resizing:")