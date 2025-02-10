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

// BoundedQueue represents a queue with a dynamic size limit and priority-based enqueueing.
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
	oldData := bq.tasks
	n := len(oldData)
	task := oldData[n-1]
	bq.tasks = oldData[0 : n-1]
	return task
}

// NewBoundedQueue creates a new BoundedQueue with the specified initial limit.
func NewBoundedQueue(limit int) *BoundedQueue {
	bq := &BoundedQueue{
		tasks: make([]PriorityTask, 0),
		limit: limit,
	}
	heap.Init(bq) // Initialize the underlying heap structure
	return bq
}

// Enqueue adds an element to the queue with a given priority.
// Returns an error if the queue is full.
func (bq *BoundedQueue) Enqueue(task interface{}, priority int) error {
	if bq.Len() >= bq.limit {
		return errors.New("queue is full")
	}
	heap.Push(bq, PriorityTask{Task: task, Priority: priority})
	return nil
}

// Dequeue removes and returns the element with the highest priority from the queue.
// Returns an error if the queue is empty.
func (bq *BoundedQueue) Dequeue() (interface{}, error) {
	if bq.Len() == 0 {
		return nil, errors.New("queue is empty")
	}
	task := heap.Pop(bq).(PriorityTask)
	return task.Task, nil
}

// Resize changes the limit of the queue, altering its capacity dynamically.
// It can accommodate upwards resizing; downwards would require handling overflow elements.
func (bq *BoundedQueue) Resize(newLimit int) {
	if newLimit < bq.Len() {
		// Optional: Define policy for what to do with existing elements
		// if trying to shrink below current number of elements.
		fmt.Println("Warning: Attempting to shrink below current element count. No elements dropped.")
	}
	bq.limit = newLimit
}

func main() {
	// Initial queue
	bq := NewBoundedQueue(2)

	// Enqueue tasks
	_ = bq.Enqueue("Task 1", 2)
	_ = bq.Enqueue("Task 2", 1)

	// Attempt to add another; should initially fail as queue is full with limit 2
	err := bq.Enqueue("Task 3", 3)
	if err != nil {
		fmt.Println(err) // Output: queue is full
	}

	// Resize the queue to add more capacity
	bq.Resize(4)
	if err := bq.Enqueue("Task 3", 3); err == nil {
		fmt.Println("Enqueued Task 3 successfully after resizing")
	}

	// Display dequeued tasks
	fmt.Println("Dequeuing tasks in priority order:")
	for bq.Len() > 0 {
		task, _ := bq.Dequeue()
		fmt.Println(task)
	}
}
