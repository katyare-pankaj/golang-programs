package main

import (
	"errors"
	"fmt"
)

// Element represents a queue element with a priority.
type Element struct {
	value    interface{}
	priority int
}

// PriorityQueue represents a priority-based fixed-size bounded queue.
type PriorityQueue struct {
	data  []Element
	limit int
}

// NewPriorityQueue creates a new PriorityQueue with the specified limit.
func NewPriorityQueue(limit int) *PriorityQueue {
	return &PriorityQueue{
		data:  make([]Element, 0),
		limit: limit,
	}
}

// Enqueue adds an element to the queue based on its priority.
func (pq *PriorityQueue) Enqueue(value interface{}, priority int) error {
	newElement := Element{value: value, priority: priority}

	// If there's room, just add the item
	if len(pq.data) < pq.limit {
		pq.insertByPriority(newElement)
		return nil
	}

	// If the queue is full, only replace if the new element has higher priority than the lowest
	if len(pq.data) == pq.limit {
		lowestPriorityIndex := pq.findLowestPriorityIndex()
		if pq.data[lowestPriorityIndex].priority < priority {
			pq.data = append(pq.data[:lowestPriorityIndex], pq.data[lowestPriorityIndex+1:]...)
			pq.insertByPriority(newElement)
			return nil
		}
	}
	return errors.New("queue is full and the new element's priority is too low")
}

// Finds the index of the element with the lowest priority
func (pq *PriorityQueue) findLowestPriorityIndex() int {
	lowestIndex := 0
	for i, elem := range pq.data {
		if elem.priority < pq.data[lowestIndex].priority {
			lowestIndex = i
		}
	}
	return lowestIndex
}

// Inserts a new element according to its priority
func (pq *PriorityQueue) insertByPriority(elem Element) {
	inserted := false
	for i, existing := range pq.data {
		if elem.priority > existing.priority {
			// Insert before the current element
			pq.data = append(pq.data[:i], append([]Element{elem}, pq.data[i:]...)...)
			inserted = true
			break
		}
	}
	if !inserted {
		// All existing elements have higher or equal priority, append to the end
		pq.data = append(pq.data, elem)
	}
}

// Dequeue removes and returns the element with the highest priority.
func (pq *PriorityQueue) Dequeue() (interface{}, error) {
	if len(pq.data) == 0 {
		return nil, errors.New("queue is empty")
	}
	element := pq.data[0]
	pq.data = pq.data[1:]
	return element.value, nil
}

// Size returns the current number of elements in the queue.
func (pq *PriorityQueue) Size() int {
	return len(pq.data)
}

// IsEmpty checks whether the queue is empty.
func (pq *PriorityQueue) IsEmpty() bool {
	return len(pq.data) == 0
}

func main() {
	// Example use: managing prioritized tasks in a queue with size limit
	queueLimit := 5
	pq := NewPriorityQueue(queueLimit)

	// Adding tasks with priorities
	tasks := []struct {
		value    string
		priority int
	}{
		{"Task 1", 5},
		{"Task 2", 3},
		{"Task 3", 6},
		{"Task 4", 2},
		{"Task 5", 4},
		{"Task 6", 7}, // Higher priority than Task 2
	}

	for _, task := range tasks {
		err := pq.Enqueue(task.value, task.priority)
		if err != nil {
			fmt.Printf("Failed to enqueue %v: %v\n", task.value, err)
		} else {
			fmt.Printf("Enqueued: %v (Priority: %d)\n", task.value, task.priority)
		}
	}

	// Dequeuing tasks
	fmt.Println("\nDequeuing tasks:")
	for !pq.IsEmpty() {
		task, _ := pq.Dequeue()
		fmt.Printf("Dequeued: %v\n", task)
	}
}
