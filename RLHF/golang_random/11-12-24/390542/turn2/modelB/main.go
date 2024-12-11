package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

type Node struct {
	value interface{}
	next  *Node
}

type LockFreeQueue struct {
	head *Node
	tail *Node
}

// NewLockFreeQueue creates a new lock-free queue.
func NewLockFreeQueue() *LockFreeQueue {
	return &LockFreeQueue{}
}

// Enqueue adds a value to the end of the queue.
func (q *LockFreeQueue) Enqueue(value interface{}) {
	newNode := &Node{value: value}
	for {
		// Get the current tail
		oldTail := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&q.tail)))

		// If the queue is empty, set both head and tail to the new node
		if oldTail == nil {
			if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&q.tail)), nil, unsafe.Pointer(newNode)) {
				atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&q.head)), unsafe.Pointer(newNode))
				return
			}
		} else {
			// Otherwise, point the new node to the current tail
			newNode.next = (*Node)(oldTail)

			// Try to update the tail to the new node
			if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&q.tail)), oldTail, unsafe.Pointer(newNode)) {
				return
			}
		}
	}
}

// Dequeue removes and returns the value from the front of the queue.
func (q *LockFreeQueue) Dequeue() (interface{}, bool) {
	for {
		oldHead := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&q.head)))
		if oldHead == nil {
			return nil, false // Queue is empty
		}

		newHead := (*Node)(oldHead).next

		// If the next node is nil, it means the queue has only one element
		if newHead == nil {
			if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&q.tail)), oldHead, nil) {
				return (*Node)(oldHead).value, true
			}
		} else {
			// Try to update the head to the next node
			if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&q.head)), oldHead, unsafe.Pointer(newHead)) {
				return (*Node)(oldHead).value, true
			}
		}
	}
}

func main() {
	queue := NewLockFreeQueue()

	// Enqueuing values
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	// Dequeuing values
	val, ok := queue.Dequeue()
	fmt.Println(val, ok) // Outputs: 1 true
	val, ok = queue.Dequeue()
	fmt.Println(val, ok) // Outputs: 2 true
	val, ok = queue.Dequeue()
	fmt.Println(val, ok) // Outputs: 3 true
	val, ok = queue.Dequeue()
	fmt.Println(val, ok) // Outputs: <nil> false (queue is empty)
}
