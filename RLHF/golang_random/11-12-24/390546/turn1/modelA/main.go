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

type LockFreeStack struct {
	top *Node
}

// Push adds a value to the top of the stack.
func (s *LockFreeStack) Push(value interface{}) {
	newNode := &Node{value: value}
	for {
		oldTop := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&s.top)))
		newNode.next = (*Node)(oldTop)

		// Try to update the top pointer to the new node.
		if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&s.top)), oldTop, unsafe.Pointer(newNode)) {
			return // Success
		}
		// If the swap failed, the top pointer has changed, go back to the top.
	}
}

// Pop removes and returns the top value from the stack.
func (s *LockFreeStack) Pop() (interface{}, bool) {
	for {
		oldTop := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&s.top)))
		if oldTop == nil {
			return nil, false // Stack is empty
		}
		newTop := (*Node)(oldTop).next

		// Try to update the top pointer to the next node.
		if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&s.top)), oldTop, unsafe.Pointer(newTop)) {
			return (*Node)(oldTop).value, true // Success
		}
		// If the swap failed, the top pointer has changed, go back to the top.
	}
}

func main() {
	stack := &LockFreeStack{}

	// Pushing values
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Popping values
	val, ok := stack.Pop()
	fmt.Println(val, ok) // Outputs: 3 true
	val, ok = stack.Pop()
	fmt.Println(val, ok) // Outputs: 2 true
	val, ok = stack.Pop()
	fmt.Println(val, ok) // Outputs: 1 true
	val, ok = stack.Pop()
	fmt.Println(val, ok) // Outputs: <nil> false (stack is empty)
}
