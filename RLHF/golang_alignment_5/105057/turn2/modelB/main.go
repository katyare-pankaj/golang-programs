package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

const (
	cacheLineSize = 64 // Adjust this based on your architecture
)

// Align ensure that the struct is aligned to the cache line size
func align(size uintptr) uintptr {
	return (size + (cacheLineSize - 1)) & ^(cacheLineSize - 1)
}

// RingBuffer is a high-performance lock-free ring buffer.
type RingBuffer struct {
	data      []interface{} // padded data slice
	writePos uint64
	readPos  uint64
	mask     uint64
	_        [(align(unsafe.Sizeof(RingBuffer{})) - unsafe.Sizeof(RingBuffer{}))]byte // padding
}

// NewRingBuffer creates a new RingBuffer with the specified size.
func NewRingBuffer(size int) *RingBuffer {
	if size <= 0 {
		panic("Ring buffer size must be positive")
	}

	// Calculate the next power of 2 for efficient masking
	size--
	size |= size >> 1
	size |= size >> 2
	size |= size >> 4
	size |= size >> 8
	size |= size >> 16
	size |= size >> 32
	size++

	dataSize := align(uintptr(unsafe.Sizeof(interface{}(nil)))) * uintptr(size)
	rb := &RingBuffer{
		data:     (*[1 << 30]interface{})(unsafe.Pointer(alignedAlloc(dataSize)))[:size],
		mask:     uint64(size - 1),
	}
	return rb
}

func alignedAlloc(size uintptr) unsafe.Pointer {
	p := make([]byte, size+cacheLineSize)
	return unsafe.Pointer(&p[0] + (cacheLineSize - (uintptr(unsafe.Pointer(&p[0])) % cacheLineSize)))
}

// WriteEvent writes an event into the ring buffer. Returns true if successful, false if the buffer is full.
func (rb *RingBuffer) WriteEvent(event interface{}) bool {
	nextWritePos := (rb.writePos + 1) & rb.mask
	if nextWritePos == atomic.LoadUint64(&rb.readPos) {
		// Buffer is full
		return false
	}

	// Atomic store to ensure consistent data update
	atomic.StoreInterface(&rb.data[rb.writePos&rb.mask], event)
	atomic.StoreUint64(&rb.writePos, nextWritePos)
	return true
}

// ReadEvent reads an event from the ring buffer. Returns true if an event was read, false if the buffer is empty.
func (rb *RingBuffer) ReadEvent() (interface{}, bool) {