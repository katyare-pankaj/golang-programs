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
	node := &Node{} // Dummy node
	return &LockFreeQueue{
		head: node,
		tail: node,
	}
}

// Enqueue adds a value to the end of the queue.
func (q *LockFreeQueue) Enqueue(value interface{}) {
	newNode := &Node{value: value}
	for {
		tail := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&q.tail)))
		next := (*Node)(tail).next

		// If the tail is pointing to the last node and it has no next, try to link the new node.
		if next == nil {
			if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&(*Node)(tail).next)), nil, unsafe.Pointer(newNode)) {
				// Successfully linked the new node, move the tail pointer forward.
				atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&q.tail)), tail, unsafe.Pointer(newNode))
				return
			}
		} else {
			// Tail was not pointing to the last node, move the tail pointer forward.
			atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&q.tail)), tail, unsafe.Pointer(next))
		}
	}
}

// Dequeue removes and returns a value from the front of the queue.
func (q *LockFreeQueue) Dequeue() (interface{}, bool) {
	for {
		head := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&q.head)))
		tail := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&q.tail)))
		next := (*Node)(head).next

		// If the queue is empty.
		if head == tail {
			if next == nil {
				return nil, false // Queue is empty
			}
			// Move the tail pointer forward if it is pointing to an outdated node.
			atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&q.tail)), head, unsafe.Pointer(next))
		} else {
			// Dequeue the front node.
			value := next.value
			if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&q.head)), head, unsafe.Pointer(next)) {
				return value, true // Successfully dequeued
			}
		}
	}
}

func main() {
	queue := NewLockFreeQueue()

	// Enqueue values
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	// Dequeue values
	val, ok := queue.Dequeue()
	fmt.Println(val, ok) // Outputs: 1 true
	val, ok = queue.Dequeue()
	fmt.Println(val, ok) // Outputs: 2 true
	val, ok = queue.Dequeue()
	fmt.Println(val, ok) // Outputs: 3 true
	val, ok = queue.Dequeue()
	fmt.Println(val, ok) // Outputs: <nil> false (queue is empty)
}
