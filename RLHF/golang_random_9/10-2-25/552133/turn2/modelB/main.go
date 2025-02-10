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

// BoundedQueue represents a queue with a maximum size limit and priority-based enqueueing.
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
}
func (bq *BoundedQueue) Pop() interface{} {
	old := bq.tasks
	n := len(old)
	task := old[n-1]
	bq.tasks = old[0 : n-1]
	return task
}

// NewBoundedQueue creates a new BoundedQueue with the specified limit.
func NewBoundedQueue(limit int) *BoundedQueue {
	return &BoundedQueue{
		tasks: make([]PriorityTask, 0),
		limit: limit,
	}
}

// Enqueue adds an element to the queue with a given priority.
// Returns an error if the queue is full.
func (bq *BoundedQueue) Enqueue(task interface{}, priority int) error {
	if len(bq.tasks) >= bq.limit {
		return errors.New("queue is full")
	}
	heap.Push(bq, PriorityTask{Task: task, Priority: priority})
	return nil
}

// Dequeue removes and returns the element with the highest priority from the queue.
// Returns an error if the queue is empty.
func (bq *BoundedQueue) Dequeue() (interface{}, error) {
	if len(bq.tasks) == 0 {
		return nil, errors.New("queue is empty")
	}
	return heap.Pop(bq).(PriorityTask).Task, nil
}

// Example usage:
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

	fmt.Println("\nBounded Queue Tests:")
	bq = NewBoundedQueue(2)
	bq.Enqueue("A", 2)
	bq.Enqueue("B", 1)
	err := bq.Enqueue("C", 3) // Queue full
	if err != nil {
		fmt.Println(err) // Output: queue is full
	}

	task, err := bq.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Dequeued:", task) // Output: Dequeued: A
	}
	bq.Enqueue("D", 2)
	task, _ = bq.Dequeue()
	fmt.Println("Dequeued:", task) // Output: Dequeued: B
}
