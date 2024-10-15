package main

import (
	"fmt"
	"sync/atomic"
)

// RingBuffer is a lock-free ring buffer implementation
type RingBuffer struct {
	data     []interface{}
	writePos uint64
	readPos  uint64
	mask     uint64
}

// NewRingBuffer creates a new ring buffer with a specific size
func NewRingBuffer(size int) *RingBuffer {
	if size <= 0 {
		return nil
	}

	nextPowerOfTwo := 1
	for nextPowerOfTwo < size {
		nextPowerOfTwo <<= 1
	}

	rb := &RingBuffer{
		data: make([]interface{}, nextPowerOfTwo),
		mask: uint64(nextPowerOfTwo - 1),
	}
	return rb
}

// Push adds an event to the ring buffer. It returns false if the buffer is full.
func (rb *RingBuffer) Push(event interface{}) bool {
	nextWritePos := (rb.writePos + 1) & rb.mask
	if nextWritePos == rb.readPos {
		// Buffer is full
		return false
	}

	rb.data[rb.writePos] = event
	atomic.StoreUint64(&rb.writePos, nextWritePos)
	return true
}

// Pop removes the oldest event from the ring buffer and returns it. It returns nil if the buffer is empty.
func (rb *RingBuffer) Pop() interface{} {
	if rb.writePos == rb.readPos {
		// Buffer is empty
		return nil
	}

	event := rb.data[rb.readPos]
	rb.data[rb.readPos] = nil
	atomic.StoreUint64(&rb.readPos, (rb.readPos+1)&rb.mask)
	return event
}

// Main function to demonstrate the usage of the RingBuffer
func main() {
	const bufferSize = 1024
	rb := NewRingBuffer(bufferSize)

	// Producer goroutine
	go func() {
		for i := 0; i < 10000; i++ {
			rb.Push(i)
		}
	}()

	// Consumer goroutine
	go func() {
		for {
			event := rb.Pop()
			if event == nil {
				break
			}
			fmt.Println("Consumed:", event)
		}
	}()
}
