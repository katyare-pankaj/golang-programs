package main

import (
	"errors"
	"fmt"
)

// BoundedQueue represents a queue with a maximum size limit.
type BoundedQueue struct {
	data  []interface{}
	limit int
}

// NewBoundedQueue creates a new BoundedQueue with the specified limit.
func NewBoundedQueue(limit int) *BoundedQueue {
	return &BoundedQueue{
		data:  make([]interface{}, 0),
		limit: limit,
	}
}

// Enqueue adds an element to the queue. Returns an error if the queue is full.
func (bq *BoundedQueue) Enqueue(element interface{}) error {
	if len(bq.data) >= bq.limit {
		return errors.New("queue is full")
	}
	bq.data = append(bq.data, element)
	return nil
}

// Dequeue removes and returns an element from the queue. Returns an error if the queue is empty.
func (bq *BoundedQueue) Dequeue() (interface{}, error) {
	if len(bq.data) == 0 {
		return nil, errors.New("queue is empty")
	}
	element := bq.data[0]
	bq.data = bq.data[1:]
	return element, nil
}

// Size returns the current size of the queue.
func (bq *BoundedQueue) Size() int {
	return len(bq.data)
}

// IsEmpty checks whether the queue is empty.
func (bq *BoundedQueue) IsEmpty() bool {
	return len(bq.data) == 0
}

// IsFull checks whether the queue is full.
func (bq *BoundedQueue) IsFull() bool {
	return len(bq.data) >= bq.limit
}

func main() {
	// Example use: managing tasks in a queue with a size limit
	queueLimit := 5
	bq := NewBoundedQueue(queueLimit)

	// Enqueuing tasks
	for i := 0; i < queueLimit+1; i++ {
		err := bq.Enqueue(fmt.Sprintf("Task %d", i+1))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Enqueued: Task %d\n", i+1)
		}
	}

	// Dequeuing tasks
	for !bq.IsEmpty() {
		task, err := bq.Dequeue()
		if err == nil {
			fmt.Printf("Dequeued: %v\n", task)
		} else {
			fmt.Println(err)
		}
	}

	fmt.Println("Finished processing tasks.")
}
