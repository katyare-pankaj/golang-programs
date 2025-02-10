package main

import (
	"errors"
	"fmt"
)

const (
	queueSize = 5 // Choose your desired fixed size
)

type FixedSizeQueue struct {
	data  []int
	front int
	rear  int
	count int
}

// NewFixedSizeQueue returns an initialized FixedSizeQueue
func NewFixedSizeQueue() *FixedSizeQueue {
	return &FixedSizeQueue{
		data: make([]int, queueSize),
	}
}

// IsEmpty checks if the queue is empty
func (q *FixedSizeQueue) IsEmpty() bool {
	return q.count == 0
}

// IsFull checks if the queue is full
func (q *FixedSizeQueue) IsFull() bool {
	return q.count == queueSize
}

// Size returns the current size of the queue
func (q *FixedSizeQueue) Size() int {
	return q.count
}

// Enqueue adds an element to the back of the queue
func (q *FixedSizeQueue) Enqueue(item int) error {
	if q.IsFull() {
		return errors.New("Queue is full")
	}
	q.data[q.rear] = item
	q.rear = (q.rear + 1) % queueSize
	q.count++
	return nil
}

// Dequeue removes and returns the element from the front of the queue
func (q *FixedSizeQueue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Queue is empty")
	}
	result := q.data[q.front]
	q.front = (q.front + 1) % queueSize
	q.count--
	return result, nil
}

func main() {
	queue := NewFixedSizeQueue()
	fmt.Println("Initial queue state:")
	queue.Display() // Output: Initial queue state: [0 0 0 0 0]
	fmt.Println()

	fmt.Println("Enqueuing elements:")
	for i := 1; i <= 7; i++ {
		err := queue.Enqueue(i)
		if err != nil {
			fmt.Println(err)
		}
	}
	queue.Display() // Output: Enqueuing elements: Queue is full
	fmt.Println()

	fmt.Println("Dequeuing elements:")
	for i := 1; i <= 3; i++ {
		element, err := queue.Dequeue()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(element)
		}
	}
	queue.Display() // Output: Dequeuing elements: 1 2 3 [4 5 0 0 0]
	fmt.Println()

	fmt.Println("Enqueuing more elements:")
	err := queue.Enqueue(8)
	if err != nil {
		fmt.Println(err)
	}
	queue.Display() // Output: Enqueuing more elements: [4 5 6 7 8]
}

// Display displays the current state of the queue
func (q *FixedSizeQueue) Display() {
	fmt.Println(q.data)
}
