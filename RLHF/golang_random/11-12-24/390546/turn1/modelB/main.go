package main

import (
	"fmt"
	"sync/atomic"
	"time"
	"unsafe"
)

type Node struct {
	value int
	next  *Node
}

type LockFreeStack struct {
	top *Node
}

func NewLockFreeStack() *LockFreeStack {
	return &LockFreeStack{}
}

func (s *LockFreeStack) Push(value int) {
	newNode := &Node{value: value}

	for {
		oldTop := atomic.LoadPointer((*unsafe.Pointer)(&s.top))
		newTop := &Node{value: value, next: (*Node)(oldTop)}

		if atomic.CompareAndSwapPointer((*unsafe.Pointer)(&s.top), oldTop, unsafe.Pointer(newTop)) {
			return
		}
	}
}

func (s *LockFreeStack) Pop() (int, bool) {
	for {
		oldTop := atomic.LoadPointer((*unsafe.Pointer)(&s.top))

		if oldTop == nil {
			return 0, false // Stack is empty
		}

		newTop := (*Node)(oldTop).next

		if atomic.CompareAndSwapPointer((*unsafe.Pointer)(&s.top), oldTop, unsafe.Pointer(newTop)) {
			return (*Node)(oldTop).value, true
		}
	}
}

func main() {
	stack := NewLockFreeStack()

	go func() {
		stack.Push(1)
		stack.Push(2)
	}()

	go func() {
		value, ok := stack.Pop()
		if ok {
			fmt.Println("Popped:", value)
		}
		value, ok = stack.Pop()
		if ok {
			fmt.Println("Popped:", value)
		}
	}()

	time.Sleep(1 * time.Second)
}
